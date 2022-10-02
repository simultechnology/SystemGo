package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("start!")

	oldFile, err := os.Open("ch3/9/english_words.tsv")
	if err != nil {
		panic(err)
	}
	newFile, err := os.Create("ch3/9/new_words.tsv")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(newFile, oldFile)
	if err != nil {
		panic(err)
	}
}
