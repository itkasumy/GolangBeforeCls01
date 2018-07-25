package main

import "fmt"

func main()  {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := arr[:]
	fmt.Println("s1 = ", s1)
	fmt.Printf("len(s1) = %d, cap(s1) = %d\n", len(s1), cap(s1))

	data := arr[0]
	fmt.Println("data = ", data)

	s2 := arr[3 : 6 : 7]
	fmt.Println("s2 = ", s2)
	fmt.Printf("len(s2) = %d, cap(s2) = %d\n", len(s2), cap(s2))

	s3 := arr[:6]
	fmt.Println("s3 = ", s3)
	fmt.Printf("len(s3) = %d, cap(s3) = %d\n", len(s3), cap(s3))

	s4 := arr[3:]
	fmt.Println("s4 = ", s4)
	fmt.Printf("len(s4) = %d, cap(s4) = %d\n", len(s4), cap(s4))
}
