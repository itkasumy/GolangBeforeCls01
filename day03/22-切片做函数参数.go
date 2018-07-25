package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitData(s []int)  {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
}

func BubbleSort(s []int)  {
	l := len(s)
	for i := 0; i < l - 1; i++ {
		for j := 0; j < l - i - 1; j++ {
			if s[j] > s[j + 1] {
				s[j], s[j + 1] = s[j + 1], s[j]
			}
		}
	}
}

func main()  {
	n := 10
	s := make([]int, n)

	InitData(s)

	fmt.Println("before sort: s = ", s)
	BubbleSort(s)
	fmt.Println("after sort: s = ", s)
}
