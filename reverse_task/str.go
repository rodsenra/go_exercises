package str
import "fmt"

func Reverse(input string) string {
        rune_array := []rune(input)
        size := len(rune_array)
	var inverted []rune
        inverted = make([]rune, size)
        fmt.Println(size)
	for i, v := range rune_array {
            fmt.Printf("%d %s\n", i, v)
            inverted[size-i-1] = v
	}
	return string(inverted)
}

