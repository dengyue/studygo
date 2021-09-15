package main

import (
	"bufio"
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
			continue
		}
		go process(conn)
	}

}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from conn failed, err", err)
			break
		}
		recStr := string(buf[:n])
		fmt.Println("receive data from client: ", recStr)
	}
}
