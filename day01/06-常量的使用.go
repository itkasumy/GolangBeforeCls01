package main

import "fmt"

func main()  {
	const a int = 10
	fmt.Println("a = ", a)
	const b = 20
	fmt.Printf("b type is %T\n", b)
	fmt.Println("b = ", b)
}