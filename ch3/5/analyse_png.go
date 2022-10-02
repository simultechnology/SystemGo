package main

import (
	funcs "./funcs"
	"os"
)

func main() {
	file, err := os.Open("ch3/5/PNG_transparency_demonstration_secret.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := funcs.ReadChunks(file)
	for _, chunk := range chunks {
		funcs.DumpChunks(chunk)
	}
}
