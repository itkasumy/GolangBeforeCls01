package main

import "fmt"

func main()  {
	a := 10
	fmt.Println("a = ", a)

	fmt.Printf("a = %d\n", a)

	b := 20
	c := 30
	fmt.Println("a = ", a, ", b = ", b, ", c = ", c)
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}
