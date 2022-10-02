package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("start!")

	header := bytes.NewBufferString("----- HEADER ------\n")
	content := bytes.NewBufferString("Example of io.MultiReader\n")
	footer := bytes.NewBufferString("----- FOOTER ------\n")

	reader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, reader)

}
