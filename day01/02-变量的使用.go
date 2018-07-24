package main

import "fmt"

func main()  {
	var a int
	fmt.Println("a = ", a)

	a = 10
	fmt.Println("a = ", a)

	var b int = 20
	b = 30
	fmt.Println("b = ", b)

	c := 40
	fmt.Printf("c type is %T\n", c)
}
