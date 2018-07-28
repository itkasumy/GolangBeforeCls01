package main

import (
	"fmt"
	"time"
)

func Priter(str string) {
	for _, data := range str{
		fmt.Print(string(data))
		time.Sleep(time.Second)
	}
}

func Person01()  {
	Priter("Hello")
}

func Person02()  {
	Priter("world")
}

func main() {
	//Priter("abcdefghijklmnopqrstuvwxyz")

	go Person01()
	go Person02()

	for {}
}
