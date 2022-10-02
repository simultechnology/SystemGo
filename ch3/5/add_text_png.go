package main

import (
	funcs "./funcs"
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"io"
	"os"
)

func main() {
	file, err := os.Open("ch3/5/PNG_transparency_demonstration_1.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	newFile, err := os.Create("ch3/5/PNG_transparency_demonstration_secret.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	chunks := funcs.ReadChunks(file)
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	io.Copy(newFile, chunks[0])
	io.Copy(newFile, textChunk("my name is Mine"))
	for _, chunk := range chunks[1:] {
		//funcs.DumpChunks(chunk)
		io.Copy(newFile, chunk)
	}
}

func textChunk(text string) *bytes.Buffer {
	byteText := []byte(text)
	crc := crc32.NewIEEE()
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, int32(len(byteText)))
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(&buffer, crc)
	io.WriteString(writer, "teXt")
	writer.Write(byteText)
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer
}
