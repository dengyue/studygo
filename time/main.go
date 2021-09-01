package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())

	//time
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	//time.Unix
	ret := time.Unix(1630415215, 0)
	fmt.Println(ret)

	//time duration
	fmt.Println(time.Second)
	fmt.Println(now.Add(time.Hour * 24))

	//Sub
	d := now.Sub(now.Add(time.Hour * 24)).Seconds()
	fmt.Printf("sub: %v \n", d)

	//time tick
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	//format time
	fmt.Println(time.Now().Format("2006-01-02"))          //2021-08-31
	fmt.Println(time.Now().Format("01/02/2006"))          //08/31/2021
	fmt.Println(time.Now().Format("2006/01/02 03:04:05")) //2021/08/31 03:04:05

	stringToTime()

	//sleep
	n := 5
	fmt.Println("start sleep......")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("5 seconds later......")
}

//convert string to time
func stringToTime() {
	//2021/08/31 09:43:32
	timeObj, err := time.Parse("2006/01/02 03:04:05", "2021/08/31 09:43:32")
	if err != nil {
		fmt.Printf("Parse time string error. err : %v \n", err)
		return
	}
	fmt.Println(timeObj)
}
