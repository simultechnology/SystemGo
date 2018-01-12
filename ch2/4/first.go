package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	message := "test is test!!\n"
	messageBytes := []byte(message)
	file, error := os.Create("test.txt")
	if error != nil {
		panic(error)
	}
	os.Stdout.Write(messageBytes)
	file.Write(messageBytes)

	var buffer bytes.Buffer
	buffer.Write(messageBytes)
	fmt.Println(buffer.String())
}
