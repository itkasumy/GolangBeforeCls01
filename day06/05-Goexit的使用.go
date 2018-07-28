package main

import (
	"fmt"
	"runtime"
)

func test()  {
	defer fmt.Println("cccccc")
	//return
	runtime.Goexit()
	fmt.Println("dddddd")
}

func main() {
	go func() {
		fmt.Println("aaaaaa")
		test()
		fmt.Println("bbbbbb")
	}()

	for {}
}
