package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("start!")

	reader := strings.NewReader("Example of io.SectionReader\n")
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
	fmt.Println("\n----------------")
	io.Copy(os.Stdout, io.NewSectionReader(reader, 15, 5))
}
