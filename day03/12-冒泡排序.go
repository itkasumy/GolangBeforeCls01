package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	var a [10]int
	n := len(a)

	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
		fmt.Printf("%d ", a[i])
	}

	fmt.Printf("\n")

	for j := 0; j < n - 1; j++ {
		for k := 0; k < n - j - 1; k++ {
			if a[k] > a[k + 1] {
				a[k], a[k + 1] = a[k + 1], a[k]
			}
		}
	}

	fmt.Println("after: a = ", a)
}
