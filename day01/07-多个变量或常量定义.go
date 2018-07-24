package main

import "fmt"

func main()  {
	//var a int
	//var b float64
	var (
		a int
		b float64
	)
	a, b = 10, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//const x int = 10
	//const y float64 = 3.14
	const (
		x int = 10
		y float64 = 3.14
	)
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
}
