package main

import "fmt"

func main()  {
	a := [5]int{}
	fmt.Printf("len(a) = %d, cap(a) = %d\n", len(a), cap(a))

	s := []int{}
	fmt.Printf("len(s) = %d, cap(s) = %d\n", len(s), cap(s))

	s = append(s, 11)
	fmt.Printf("append: len(s) = %d, cap(s) = %d\n", len(s), cap(s))

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println("s1 = ", s1)

	s2 := make([]int, 5, 10)
	fmt.Printf("len(s2) = %d, cap(s2) = %d\n", len(s2), cap(s2))

	s3 := make([]int, 5)
	fmt.Printf("len(s3) = %d, cap(s3) = %d\n", len(s3), cap(s3))
}
