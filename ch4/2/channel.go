package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub()")
	done := make(chan bool)
	go func() {
		done <- true
		for i := 0; i < 100; i++ {
			fmt.Printf("%d.sub() is finished\n", i+1)
		}
	}()
	time.Sleep(1 * time.Nanosecond)
	<-done
	fmt.Println("all tasks are finished")
}
