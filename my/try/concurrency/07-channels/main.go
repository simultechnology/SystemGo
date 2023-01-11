package main

import (
	"log"
)

func main() {
	counterChan := make(chan int)
	go func() {
		for i := 0; i <= 32; i++ {
			counterChan <- 1
			// time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i <= 32; i++ {
			counterChan <- 1
			// time.Sleep(100 * time.Millisecond)
		}
	}()

	{
		var counter = 0
		for counter < 64 {
			counter += <-counterChan
			log.Println(counter)
		}
	}
}
