package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "World",
	}
	for i := 0; i < 100; i++ {
		source[fmt.Sprintf("%s%d", "Hello", i)] = "World"
	}
	sourceBytes, err := json.Marshal(source)
	if err != nil {
		panic(err)
	}
	gzWriter := gzip.NewWriter(w)
	gzWriter.Header.Name = "test.txt"
	defer gzWriter.Close()

	file, err := os.Create("./test.txt.log")
	writer := io.MultiWriter(gzWriter, os.Stdout, file)
	_, err = writer.Write(sourceBytes)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":55580", nil)
	if err != nil {
		return
	}
}
