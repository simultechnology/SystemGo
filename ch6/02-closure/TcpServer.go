package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"runtime"
	"strings"
)

func main() {
	fmt.Println(runtime.NumCPU(), runtime.GOMAXPROCS(-1))
	ln, err := net.Listen("tcp", "[::1]:6060")
	if err != nil {
		panic(err)
	}
	count := 0
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go func(thisCount int) {
			count += 1
			fmt.Printf("a request!!! %d\n", count)

			// fmt.Printf("Accept %v\n", conn.RemoteAddr()) // read a request
			_, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}
			//dump, err := httputil.DumpRequest(request, true)
			//if err != nil {
			//	panic(err)
			//}
			//fmt.Println(string(dump))

			header := make(map[string][]string)
			header["Access-Control-Allow-Origin"] = []string{"*"}
			header["Content-Type"] = []string{"application/json; charset=UTF-8"}

			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Header:     header,
				Body: io.NopCloser(
					strings.NewReader("Hello!!!")),
			}
			response.Write(conn)
			conn.Close()
		}(count)
	}
}
