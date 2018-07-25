package main

import "fmt"

func main()  {
	s1 := []int{}
	fmt.Println("s1 = ", s1)
	fmt.Printf("len(s1) = %d, cap(s1) = %d\n", len(s1), cap(s1))

	s1 = append(s1, 111)
	s1 = append(s1, 222)
	s1 = append(s1, 333)
	fmt.Println("s1 = ", s1)
	fmt.Printf("len(s1) = %d, cap(s1) = %d\n", len(s1), cap(s1))

	s2 := []int{1, 2, 3}
	fmt.Println("s2 = ", s2)
	s2 = append(s2, 666)
	s2 = append(s2, 777)
	s2 = append(s2, 888)
	fmt.Println("s2 = ", s2)
}
