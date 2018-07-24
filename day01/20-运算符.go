package main

import "fmt"

func main()  {
	fmt.Println("4 > 3的结果是:", 4 > 3)
	fmt.Println("4 != 3的结果是:", 4 != 3)

	fmt.Println("!(4 > 3)的结果是:", !(4 > 3))
	fmt.Println("!(4 != 3)的结果是:", !(4 != 3))

	fmt.Println("true && true 的结果是:", true && true)
	fmt.Println("true && false 的结果是:", true && false)

	fmt.Println("true || false 的结果是:", true || false)
	fmt.Println("false || false 的结果是:", false || false)

	a := 8
	fmt.Println("0 <= a && a <= 10 的结果是:", 0 <= a && a <= 10)
}
