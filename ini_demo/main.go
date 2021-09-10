package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type MySQLConfig struct {
	Address  string `ini:"address`
	Port     int    `ini:"port`
	Username string `ini:"username`
	Password string `ini:"password`
}

func loadIni(v interface{}) {
	file, err := os.Open("./config.ini")
	if err != nil {
		fmt.Printf("open config.ini failed, error: %v\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(string(line))
			}
			break
		}
		if err != nil {
			fmt.Printf("read config.ini failed, error: %v\n", err)
			break
		}
		fmt.Println(string(line))
	}
}

func main() {
	var config MySQLConfig
	loadIni(&config)
	fmt.Println(config.Address, config.Port, config.Username, config.Password)
}
