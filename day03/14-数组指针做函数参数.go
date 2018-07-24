package main

import "fmt"

func modify(p *[5]int) {
	(*p)[0] = 666
	fmt.Println("modify: *p = ", *p)
}

func main()  {
	a := [5]int{1, 2, 3, 4, 5}

	modify(&a)
	fmt.Println("main: a = ", a)
}
