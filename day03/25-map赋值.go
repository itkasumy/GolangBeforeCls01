package main

import "fmt"

func main()  {
	m := map[int]string{1: "Python", 2: "Java"}
	fmt.Println("m = ", m)
	m[1] = "PHP"
	fmt.Println("m = ", m)
	m[3] = "Ruby"
	fmt.Println("m = ", m)
}
