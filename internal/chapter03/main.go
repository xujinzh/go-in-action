package main

import (
	"fmt"
	"os"

	"github.com/xujinzh/go-in-action/internal/chapter03/words"
)

func main() {
	filename := os.Args[1]

	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)
	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}
