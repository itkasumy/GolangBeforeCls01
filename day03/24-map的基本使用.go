package main

import "fmt"

func main()  {
	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	fmt.Println("len(m1)", len(m1))

	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Println("len(m2)", len(m2))

	m3 := make(map[int]string, 2)
	m3[1] = "Java"
	m3[2] = "Golang"
	m3[3] = "Python"
	m3[4] = "JavaScript"
	fmt.Println("m3 = ", m3)
	fmt.Println("len(m3)", len(m3))

	m4 := map[int]string{1: "C", 2: "C#", 3: "Kotlin"}
	fmt.Println("m4 = ", m4)
}
