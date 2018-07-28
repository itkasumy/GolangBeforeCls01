package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func Priter(str string) {
	for _, data := range str{
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func Person01()  {
	Priter("Hello")
	ch <- 666
}

func Person02()  {
	<-ch
	Priter("world")
}

func main() {
	//Priter("abcdefghijklmnopqrstuvwxyz")

	go Person01()
	go Person02()

	for {}
}
