package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v\n", time.Now())

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})
	request, err := http.NewRequest("GET", "https://simultechnology.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "added to header")
	request.Write(os.Stdout)
}
