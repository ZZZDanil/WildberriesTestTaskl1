package main

import (
	"fmt"
	"strings"
)

func main() {
	// разбиение по словам
	// движение по массиву и
	// перестановка противоположных элементов

	text := "snow dog sun"
	words := strings.Split(text, " ")
	l := len(words) - 1
	sdrow := make([]string, l+1)
	for i := range words {
		sdrow[i] = words[l-i]
	}
	fmt.Println(text)
	fmt.Println(sdrow)

}
