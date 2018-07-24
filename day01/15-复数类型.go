package main

import "fmt"

func main()  {
	var t complex128
	t = 2.2 + 3.14i
	fmt.Println("t = ", t)

	x := 2 + 3.14i
	fmt.Printf("type of x is %T\n", x)

	fmt.Println("real(x) = ", real(x), "imag(x) = ", imag(x))
}
