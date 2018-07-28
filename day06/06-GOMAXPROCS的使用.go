package main

import (
	"runtime"
	"fmt"
)

func main() {
	n := runtime.GOMAXPROCS(8)
	fmt.Println("n = ", n)

	for {
		go fmt.Print(1)
		fmt.Print(0)
	}
}
