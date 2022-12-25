package main

import (
	"log"
	"sync"
	"time"
)

var counter = 0
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		for i := 0; i <= 32; i++ {
			counter++
			time.Sleep(100 * time.Millisecond)
			log.Println("Thread 1: ", counter)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i <= 32; i++ {
			counter++
			time.Sleep(100 * time.Millisecond)
			log.Println("Thread 2: ", counter)
		}
		wg.Done()
	}()

	wg.Wait()
}
