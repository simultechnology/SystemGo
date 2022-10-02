package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	message := "Test is test!!\n"
	messageBytes := []byte(message)
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(messageBytes)
	if err != nil {
		return
	}
	_, err = file.Write(messageBytes)
	if err != nil {
		return
	}

	var buffer bytes.Buffer
	buffer.Write(messageBytes)
	fmt.Println(buffer.String())
}
