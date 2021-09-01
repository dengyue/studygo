package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open failed")
		return
	}
	defer file.Close()

	//read file
	var tmp = make([]byte, 1024)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("read finished")
			return
		}
		if err != nil {
			fmt.Printf("read from the file failed，err %v", err)
			return
		}
		fmt.Printf("read %d byte \n", n)
		fmt.Println(string(tmp[:n]))
		if n < 1024 {
			return
		}
	}
}
