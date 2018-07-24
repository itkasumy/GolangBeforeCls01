package main

import "fmt"

func test() (a, b, c int) {
	return 1, 2, 3
}

func main()  {
	a, b := 10, 20
	var temp int
	temp = a
	a = b
	b = temp
	fmt.Printf("a = %d, b = %d\n", a, b)

	c, d := 10, 20
	c, d = d, c
	fmt.Printf("c = %d, d = %d\n", c, d)

	temp, _ = c, d
	fmt.Println("temp = ", temp)

	_, e, _ := test()
	fmt.Println("e = ", e)
}
