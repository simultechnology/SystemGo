package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/big"
	"os"
)

func main() {
	fmt.Println("start!")
	file, err := os.Create("ch3/9/new_binary_file")
	if err != nil {
		panic(err)
	}

	var buffer []byte

	// 0 以上 n 未満のセキュアなランダム整数を生成
	var n int64 = 1_000_000_000

	for i := 0; i < 1000; i++ {
		val, err := rand.Int(rand.Reader, big.NewInt(n))
		if err != nil {
			panic(err)
		}
		bytes := val.Bytes()
		buffer = append(buffer, bytes...)
	}
	fmt.Printf("%v", buffer)
	// data := []byte{0x0, 0x0, 0x27, 0x10, 0x10}
	//_, err = rand.Read(data)
	//if err != nil {
	//	panic(err)
	//}

	err = binary.Write(file, binary.BigEndian, buffer)
	if err != nil {
		panic(err)
	}
}
