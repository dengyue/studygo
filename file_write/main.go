package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	writeFile1()
	writeFile2()
	writeFile3()
}

func writeFile1() {
	fileObj, err := os.OpenFile("./tt.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("read file failed")
		return
	}
	defer fileObj.Close()

	fileObj.Write([]byte("hello, this is the demo file content \n"))
	fileObj.WriteString("write string directly. ")

}

func writeFile2() {
	fileObj, err := os.OpenFile("./tt_1.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("read file failed")
		return
	}
	defer fileObj.Close()

	wr := bufio.NewWriter(fileObj)
	wr.WriteString("write string directly. ")
	wr.Flush()

}

func writeFile3() {
	hello := "write string directly. "
	err := ioutil.WriteFile("./tt_2.txt", []byte(hello), 0664)
	if err != nil {
		fmt.Println("read file failed")
		return
	}
}
