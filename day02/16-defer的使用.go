package main

import "fmt"

func main()  {
	defer fmt.Println("666")
	fmt.Println("333")
}
