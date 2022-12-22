package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub()")
	done := make(chan bool)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("%d.sub() is finished\n", i+1)
			time.Sleep(50 * time.Millisecond)
		}
		done <- true
	}()
	fmt.Println("start a channel")
	// time.Sleep(1 * time.Nanosecond)
	res, ok := <-done
	fmt.Println(res, ok)
	fmt.Println("all tasks are finished")
}
