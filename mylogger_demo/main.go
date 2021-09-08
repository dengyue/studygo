package main

import (
	"time"

	"dengyue.org/studygo/mylogger"
)

func main() {

	log := mylogger.NewLog("debug")

	for {
		log.Debug("this is debug")
		log.Info("this is info log")
		log.Error("this error log")
		time.Sleep(time.Second)
	}
}
