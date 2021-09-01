package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	_, err := copyFile("dst.txt", "src.txt")
	if err != nil {
		fmt.Printf("copy file failed, %v", err)
		return
	}
	fmt.Printf("copy finished.")
}

func copyFile(dstFile, srcFile string) (int64, error) {
	//read file
	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Printf("read the %s error,  %v \n", srcFile, err)
		return 0, err
	}
	defer src.Close()

	//write file
	dst, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		fmt.Printf("write file %s error, %v \n ", dstFile, err)
		return 0, err
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
