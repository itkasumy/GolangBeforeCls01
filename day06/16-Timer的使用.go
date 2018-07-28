package main

import (
	"time"
	"fmt"
)

func main() {
	//timer := time.NewTimer(time.Second * 2)
	//fmt.Println("当前时间是：", time.Now())
	//
	//t := <-timer.C
	//fmt.Println("t = ", t)

	timer := time.NewTimer(time.Second)

	for {
		<-timer.C
		fmt.Println("时间到...")
	}
}
