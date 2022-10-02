package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("ch3/4/sample.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		return
	}
}
