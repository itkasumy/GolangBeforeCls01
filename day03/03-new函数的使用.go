package main

import "fmt"

func main()  {
	var p *int
	p = new(int)
	*p = 666
	fmt.Println("*p = ", *p)

	q := new(int)
	*q = 777
	fmt.Println("*q = ", *q)
}
