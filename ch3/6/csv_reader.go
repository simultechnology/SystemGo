package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("ch3/6/13TOKYO.CSV")
	if err != nil {
		panic(err)
	}
	csvReader := csv.NewReader(file)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line[2], line[6:9])
	}
}
