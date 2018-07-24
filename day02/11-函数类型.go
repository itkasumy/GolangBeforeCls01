package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

type fnType func(int, int) int

func main()  {
	var res int
	res = Add(1, 2)
	fmt.Println("res = ", res)

	var fnTest fnType
	fnTest = Add
	res = fnTest(2, 3)
	fmt.Println("res = ", res)

	fnTest = Minus
	res = fnTest(777, 666)
	fmt.Println("res = ", res)
}
