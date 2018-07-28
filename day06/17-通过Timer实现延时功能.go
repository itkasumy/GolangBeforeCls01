package main

import (
	"fmt"
	"time"
)

func main() {
	//timer := time.NewTimer(time.Second * 2)
	//<-timer.C

	//time.Sleep(time.Second * 2)

	<-time.After(time.Second * 2)

	fmt.Println("时间到")
}
