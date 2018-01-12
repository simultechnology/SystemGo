package main

import (
	"./mysort"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := ioutil.ReadFile("./my/try/uniq/numbers.txt")
	if err != nil {
		panic(err)
	}
	imports := make(map[string]struct{})
	lines := strings.Split(string(fileBytes), "\n")
	for _, line := range lines {
		imports[line] = struct{}{}
	}
	var numbers []int
	for key, _ := range imports {
		if key != "" {
			atoi, _ := strconv.Atoi(key)
			numbers = append(numbers, atoi)
		}
	}
	mysort.BubbleSort(numbers)
	var buf string
	for _, num := range numbers {

		buf += fmt.Sprintf("%v\n", num)
	}
	fmt.Printf("%v\n", buf)
	ioutil.WriteFile("./my/try/uniq/output.txt", ([]byte)(buf), os.ModePerm)
}
