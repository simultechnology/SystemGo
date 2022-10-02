package funcs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func DumpChunks(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
	if bytes.Equal(buffer, []byte("teXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
}

func ReadChunks(file *os.File) []io.Reader {
	var chunks []io.Reader
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chunks = append(chunks,
			io.NewSectionReader(file, offset, int64(length)+12))
		offset, _ = file.Seek(int64(length+8), 1)

	}
	return chunks
}
