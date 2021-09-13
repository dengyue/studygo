package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp server error.")
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed error, err:", err)
			return
		}
		for {
			var tmp [128]byte
			n, err := conn.Read(tmp[:])
			if err != nil {
				fmt.Println("read from conn failed, err", err)
				return
			}
			fmt.Println(string(tmp[:n]))
		}

	}

}
