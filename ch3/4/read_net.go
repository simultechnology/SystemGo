package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "www.api.org:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: www.api.org\r\n\r\n"))
	response, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Header)
	fmt.Println("---------------------------------------------------------")
	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body)

}
