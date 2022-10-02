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
	_, err = conn.Write([]byte("GET / HTTP/1.0\r\nHost: www.api.org\r\n\r\n"))
	if err != nil {
		panic(err)
	}
	//_, err = io.Copy(os.Stdout, conn)
	//if err != nil {
	//	panic(err)
	//}
	response, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Header)
	fmt.Println("---------------------------------------------------------")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		panic(err)
	}

}
