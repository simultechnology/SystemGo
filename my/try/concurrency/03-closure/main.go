package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()
	var counter = 0
	for i := 0; i <= 65536; i++ {
		counter++
		go func(thisCounter int) {
			log.Println("Start", thisCounter)
		}(counter)
	}
	elapsed := time.Since(start)
	log.Println(elapsed)
}
