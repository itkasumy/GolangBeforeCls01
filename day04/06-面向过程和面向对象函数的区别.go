package main

import "fmt"

func Add01(a, b int) int {
	return a + b
}

type long int

func (temp long) Add02(other long) long {
	return temp + other
}

func main() {
	var result int
	result = Add01(1, 1)
	fmt.Println("result = ", result)

	var a long = 2
	res := a.Add02(3)
	fmt.Println("res = ", res)
}
