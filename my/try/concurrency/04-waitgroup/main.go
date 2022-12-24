package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var counter = 0
	var wg sync.WaitGroup

	for i := 0; i <= 32; i++ {
		wg.Add(1)
		counter++
		go func(thisCounter int) {
			log.Println("Start", thisCounter)
			wg.Done()
		}(counter)
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Println(elapsed)
}
