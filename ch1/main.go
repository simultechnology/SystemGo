package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	name := string([]byte{72, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 33, 10})
	fmt.Println(name)
}
