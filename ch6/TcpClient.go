package main

import "net"

func main() {
	_, err := net.Dial("tcp", "[::1]:6060")
	if err != nil {
		panic(err)
	}
}
