package main

import "fmt"

var a int = 10

func fn()  {
	fmt.Println("fn: a = ", a)
}

func main()  {
	fmt.Println("main: a = ", a)
	a = 100
	fn()
}