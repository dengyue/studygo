package main

import (
	"fmt"
	"os"
)

func main() {
	insertString()
}

func insertString() {
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("read and write failed")
		return
	}
	defer fileObj.Close()

	tempFile, err := os.OpenFile("./sb.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("create temp file failed, error: %v\n", err)
		return
	}
	defer tempFile.Close()

	fileObj.Seek(1, 0)

	var s = []byte{'c'}
	fileObj.Write(s)

	// var ret [1]byte
	// n, err := fileObj.Read(ret[:])
	// if err != nil {
	// 	fmt.Printf("read from file failed. err: %v", err)
	// 	return
	// }
	// fmt.Println(string(ret[:n]))
}
