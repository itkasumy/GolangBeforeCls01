package main

import "fmt"

type Student struct {
	id int
	name string
	age int
	sex byte
	addr string
}

func main()  {
	var s1 Student = Student{1, "ksm", 18, 'n', "xiamen"}
	fmt.Println("s1 = ", s1)

	s2 := Student{id: 666, name: "ksm"}
	fmt.Println("s2 = ", s2)
}
