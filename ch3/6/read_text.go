package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var source = `first line
second line
third line
`

type Pair struct {
	Key   string
	Value int
}
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%v\n", line)
		if err == io.EOF {
			break
		}
	}
	scaner := bufio.NewScanner(strings.NewReader(source))
	for scaner.Scan() {
		fmt.Printf("%v\n", scaner.Text())
	}
	file, err := os.Open("ch3/6/english.txt")
	if err != nil {
		panic(err)
	}
	counter := make(map[string]int)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		text = strings.Trim(text, ",")
		text = strings.Trim(text, ".")
		text = strings.Trim(text, "\"")
		text = strings.Trim(text, "?")
		if text == "" {
			continue
		}
		fmt.Printf("%v ", text)
		if counter[text] == 0 {
			counter[text] = 1
		} else {
			counter[text] += 1
		}
	}
	fmt.Printf("\n%v\n", counter)

	pairList := make(PairList, len(counter))
	count := 0
	for key, value := range counter {
		pairList[count] = Pair{Key: key, Value: value}
		count += 1
	}
	sort.Sort(sort.Reverse(pairList))
	for _, value := range pairList {
		fmt.Printf("%v : %v\n", value.Key, value.Value)
	}
}
