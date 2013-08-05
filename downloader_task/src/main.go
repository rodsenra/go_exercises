package main

import (
	"bytes"
	"code.google.com/p/go.net/html"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type configuration_type struct {
	download_folder string
	workers         int
	target_url      string
}

func loadConfig() configuration_type {
	var config configuration_type
	flag.StringVar(
		&config.download_folder,
		"d", "/tmp",
		"define download folder where to store retrieved files")
	flag.StringVar(
		&config.target_url,
		"u", "",
		"define the target URL of the page containing links to download")
	flag.IntVar(
		&config.workers,
		"w", 10,
		"define the number of workers to retreive links in concurrently")
	flag.Parse()
	return config
}

func fetcher(url string, finish chan<- bool) chan string {
	page_chan := make(chan string)
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Failed to retrieve %v with error %v\n", url, err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed to read the body of %v with error %v\n", url, err)
		}
		page_chan <- string(body)
		fmt.Println("Fetcher is done")
		finish <- true
	}()
	return page_chan
}

func linkParser(page_chan chan string) <-chan string {
	link_chan := make(chan string)
	go func() {
		for page := range page_chan {
			//page := <-page_chan
			page_bytes := bytes.NewBufferString(page)
			d := html.NewTokenizer(io.Reader(page_bytes))
			for {
				tokenType := d.Next()
				if tokenType == html.ErrorToken {
					fmt.Println("\nFinished to parse page")
					break
				}
				token := d.Token()
				switch tokenType {
				case html.StartTagToken:
					if strings.EqualFold(token.Data, "A") {
						for _, a := range token.Attr {
							if strings.EqualFold(a.Key, "HREF") {
								link_chan <- a.Val
							}
						}
					}
				}
			}
		}
		close(link_chan)
	}()
	return link_chan
}

func downloader(link_chan <-chan string, done chan bool) {
	go func() {
		for link := range link_chan {
			fmt.Printf("Fetching: %v\n", link)
		}
		done <- true
	}()
}

func waitAll(done_chan chan bool, worker_count int) {
	waiting := worker_count
	for done := range done_chan {
		if done {
			waiting -= 1
		}
		if waiting == 0 {
			break
		}
	}
}

func main() {
	config := loadConfig()
	fmt.Printf("\nFetch links from %v saving in %v (concurrency:%d)\n",
		config.target_url,
		config.download_folder,
		config.workers)

	finish_fetcher := make(chan bool)
	page_chan := fetcher(config.target_url, finish_fetcher)
	link_chan := linkParser(page_chan)

	done_downloaders := make(chan bool)
	for i := 0; i < config.workers; i++ {
		downloader(link_chan, done_downloaders)
	}
	<-finish_fetcher
	// All fetchers terminated, notify parser to finish
	close(page_chan)
	waitAll(done_downloaders, config.workers)
}
