package main

import "fmt"

func main()  {
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else {
		fmt.Println("a < 10")
	}

	b := 10
	if b == 10 {
		fmt.Println("b == 10")
	}
	if b > 10 {
		fmt.Println("b > 10")
	}
	if b < 10 {
		fmt.Println("b < 10")
	}
}
