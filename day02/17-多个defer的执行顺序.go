package main

import "fmt"

func fn(a int) (res int) {
	res = 100 / a
	return
}

func main()  {
	defer fmt.Println("aaa")
	defer fmt.Println("bbb")
	defer fmt.Println(fn(0))
	defer fmt.Println("ddd")
}
