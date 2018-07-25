package main

import "fmt"

func main()  {
	m := map[int]string{
		1: "JavaScript",
		2: "Golang",
		3: "Python",
		4: "Java",
		5: "C/C++/C#",
	}
	for key, val := range m{
		fmt.Printf("key: %d => val: %s\n", key, val)
	}

	val, ok := m[0]
	if ok {
		fmt.Println("m[1] = ", val)
	} else {
		fmt.Printf("key不存在\n")
	}
}
