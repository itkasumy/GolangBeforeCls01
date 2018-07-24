package main

import "fmt"

func fn() int {
	return 666
}

func fn02() (res int) {
	res = 777
	return
}

func main()  {
	var a int
	a = fn()
	fmt.Println("a = ", a)

	b := fn02()
	fmt.Println("b = ", b)
}
