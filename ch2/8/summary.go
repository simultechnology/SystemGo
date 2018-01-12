package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{"a", "b", "c"})
	writer.Flush()

	file, err := os.Create("csv.txt")
	if err != nil {
		panic(err)
	}
	fileWriter := csv.NewWriter(file)
	fileWriter.Write([]string{"ishikawa", "foot", "ball"})
	fileWriter.Flush()
}
