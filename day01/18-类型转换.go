package main

import "fmt"

func main()  {
	var flag bool
	flag = true
	fmt.Printf("flag = %t\n", flag)

	var ch byte
	ch = 'a'
	var i int
	i = int(ch)
	fmt.Println("i = ", i)
}
