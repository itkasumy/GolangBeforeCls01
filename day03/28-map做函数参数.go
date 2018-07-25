package main

import "fmt"

func DeleteMap(m map[int]string)  {
	delete(m, 6)
}

func main()  {
	m := map[int]string{
		1: "JavaScript",
		2: "Golang",
		3: "Python",
		4: "Java",
		5: "C/C++/C#",
		6: "PHP",
	}
	fmt.Println("m = ", m)

	DeleteMap(m)
	fmt.Println("m = ", m)
}
