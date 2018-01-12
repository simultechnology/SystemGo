package main

import (
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "simultechnology.com:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: simultechnology.com\r\n\r\n")
	io.Copy(os.Stdout, conn)

	conn2, err := net.Dial("tcp", "simultechnology.com:80")
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("GET", "https://simultechnology.com", nil)
	if err != nil {
		panic(err)
	}
	request.Write(conn2)
	io.Copy(os.Stdout, conn2)
}
