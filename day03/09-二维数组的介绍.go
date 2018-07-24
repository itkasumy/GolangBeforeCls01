package main

import "fmt"

func main()  {
	var a [3][4]int
	k := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d] = %d; ", i, j, k)
		}
		fmt.Println("")
	}

	b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println("b = ", b)

	c := [3][4]int{{1, 2, 3}, {4, 5, 6}, {7, 8}}
	fmt.Println("c = ", c)
}
