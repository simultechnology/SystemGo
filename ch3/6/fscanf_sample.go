package main

import (
	"fmt"
	"strings"
)

var stringSource = "123 1.234 1.0e4 test"

func main() {
	reader := strings.NewReader(stringSource)
	var i int
	var f, g float64
	var s string
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%v f=%v g=%v s=%s", i, f, g, s)
}
