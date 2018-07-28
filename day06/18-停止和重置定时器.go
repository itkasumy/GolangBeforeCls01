package main

import (
	"time"
	"fmt"
)

func main() {
	//timer := time.NewTimer(time.Second * 3)
	//go func() {
	//	<-timer.C
	//	fmt.Println("子协程时间到，可以打印了")
	//}()
	//
	//timer.Stop()
	//
	//for {}

	timer := time.NewTimer(time.Second * 10)
	timer.Reset(time.Second)
	<-timer.C
	fmt.Println("时间到")
}
