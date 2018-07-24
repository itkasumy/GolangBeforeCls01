package main

import "fmt"

func MyFn01(a int) {
	fmt.Println("a = ", a)
}

func MyFn02(a int, b int) {
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func MyFn03(a, b int) {
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func MyFn04(a int, b float64, c string) {
	fmt.Printf("a = %d, b = %f, c = %s\n", a, b, c)
}

func main()  {
	MyFn01(666)
	MyFn02(666, 777)
	MyFn03(666, 777)
	MyFn04(666, 777.7777, "abcdef")
}
