package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start!")
	timer := time.After(10 * time.Second)
	<-timer
	fmt.Println("end!")
	//select {
	//case <-time.After(10 * time.Second):
	//	fmt.Println("end!")
	//}
}
