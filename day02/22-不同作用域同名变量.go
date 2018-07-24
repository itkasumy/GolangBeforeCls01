package main

import "fmt"

var a float32

func fn()  {
	fmt.Printf("%T\n", a)
}

func main()  {
	var a int
	fmt.Printf("%T\n", a)

	{
		var a float64
		fmt.Printf("%T\n", a)
	}

	fn()
}
