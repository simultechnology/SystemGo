package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	_, err := os.Stdout.Write([]byte("about to start!!\n"))
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(os.Stdout, "start!\n")
	if err != nil {
		return
	}
	_, err = io.WriteString(os.Stdout, "has started!\n")
	if err != nil {
		return
	}
	writer := csv.NewWriter(os.Stdout)
	err = writer.Write([]string{"a", "b", "c"})
	if err != nil {
		return
	}
	writer.Flush()

	file, err := os.Create("./ch2/8/csv.txt")
	if err != nil {
		panic(err)
	}
	fileWriter := csv.NewWriter(file)
	err = fileWriter.Write([]string{"ishikawa", "foot", "ball", "lover"})
	if err != nil {
		return
	}
	fileWriter.Flush()
}
