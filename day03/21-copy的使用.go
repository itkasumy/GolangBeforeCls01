package main

import "fmt"

func main()  {
	srcSlice := []int{1, 2}
	tarSlice := []int{6, 6, 6, 6, 6}
	copy(tarSlice, srcSlice)
	fmt.Println("tarSlice = ", tarSlice)
}
