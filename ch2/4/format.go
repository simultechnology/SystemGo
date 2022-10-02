package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	_, err2 := fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v\n", time.Now())
	if err2 != nil {
		return
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	err := encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})
	if err != nil {
		return
	}
	request, err := http.NewRequest("GET", "https://simultechnology.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "added to header")
	err = request.Write(os.Stdout)
	if err != nil {
		return
	}
}
