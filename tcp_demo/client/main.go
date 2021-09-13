package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dail failed")
		return

	}
	conn.Write([]byte("hello world!"))
	conn.Close()
}
