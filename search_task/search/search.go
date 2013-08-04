package search

import "strings"
import "io/ioutil"

type Index struct {
	files []string
}

func NewIndex(files ...string) (*Index, error) {
	var store Index
	for _, item_file := range files {
		store.files = append(store.files, item_file)
	}
	return &store, nil
}

func (i *Index) Search(term string) (found_in_files []string, err error) {

	for _, item_file := range i.files {
		var raw []byte
		raw, err = ioutil.ReadFile(item_file)
		if err != nil {
			return
		}
		text := string(raw)
		if strings.Contains(text, term) {
			found_in_files = append(found_in_files, item_file)
		}
	}

	return found_in_files, nil
}
