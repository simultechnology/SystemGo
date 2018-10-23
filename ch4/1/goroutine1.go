package main

import (
	"fmt"
	"time"
)

func sub() {
	fmt.Println("sub() is running")
	time.Sleep(2 * time.Millisecond)
	fmt.Printf("\nsub() is finished\n")
}

func main() {
	fmt.Println("start sub()")
	go sub()

	go func() {
		for {
			fmt.Printf("1")
		}
	}()

	time.Sleep(5 * time.Millisecond)
}
