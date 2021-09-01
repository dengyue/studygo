package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	useBufio()
}

func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Printf("your input : %s \n", s)
}
