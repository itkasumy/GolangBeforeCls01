package main

import "fmt"

func fn() int {
	var x int
	x++
	return x * x
}

func fn2() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main()  {
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())

	fmt.Println("=================")

	f := fn2()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
