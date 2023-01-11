package main

import (
	"log"
	"sync"
	"time"
)

var Tasks []string

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	run("T1", &wg)
	run("T2", &wg)

	wg.Wait()
}

func run(id string, wg *sync.WaitGroup) {
	go func() {
		time.Sleep(1 * time.Second)
		Tasks = append(Tasks, id)
		log.Println(Tasks)
		wg.Done()
	}()
}
