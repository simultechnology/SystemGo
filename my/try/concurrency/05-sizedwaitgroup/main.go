package main

import (
	"log"
	"time"

	"github.com/remeh/sizedwaitgroup"
)

func main() {
	start := time.Now()
	var counter = 0
	var swg = sizedwaitgroup.New(6)

	for i := 0; i <= 32; i++ {
		swg.Add()
		counter++
		go func(thisCounter int) {
			log.Println("Start", thisCounter)
			time.Sleep(2 * time.Second)
			swg.Done()
		}(counter)
	}
	swg.Wait()
	elapsed := time.Since(start)
	log.Println(elapsed)
}
