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
	var s Student
	s.id = 123456
	s.name = "ksm"
	s.age = 18
	s.sex = 'g'
	s.addr = "shanghai"
	fmt.Println("s = ", s)
}
