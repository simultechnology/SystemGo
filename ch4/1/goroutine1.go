package main

import (
	"fmt"
	"time"
)

func sub() {
	fmt.Println("sub() is running")
	time.Sleep(2 * time.Second)
	fmt.Printf("\nsub() is finished\n")
}

func main() {
	fmt.Println("start sub()")
	go sub()

	go func() {
		for {
			fmt.Printf("1")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(10 * time.Second)
}
