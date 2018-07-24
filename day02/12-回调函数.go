package main

import "fmt"

type fnType func(int, int) int

func Add(a, b int) int {
	return a + b
}

func Calc(a, b int, fn fnType) (res int) {
	fmt.Println("Calc")
	res = fn(a, b)
	return
}

func main()  {
	res := Calc(1, 2, Add)
	fmt.Println("res = ", res)
}
