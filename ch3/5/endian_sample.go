package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	fmt.Println("start!")

	data := []byte{0x0, 0x0, 0x27, 0x10}

	var i int32
	err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &i)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data: %d\n", i)

	strData := []byte{72, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 33, 10}
	var s string
	err = binary.Read(bytes.NewReader(strData), binary.LittleEndian, &s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data: %d\n", s)
}
