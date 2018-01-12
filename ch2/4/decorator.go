package main

import (
	"bufio"
	"compress/gzip"
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	//writer.Write([]byte("io.MultiWriter sample!"))
	io.WriteString(writer, "io.MultiWriter sample 2!!!\n")

	gzFile, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	gzWriter := gzip.NewWriter(gzFile)
	gzWriter.Header.Name = "test.txt"
	io.WriteString(gzWriter, "gzip.Writer example\n")
	gzWriter.Close()

	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}
