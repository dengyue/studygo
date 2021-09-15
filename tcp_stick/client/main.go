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
	defer conn.Close()
	for i := 0; i < 100; i++ {
		msg := `hello, hello. How are you?`
		conn.Write([]byte(msg))
	}
}
