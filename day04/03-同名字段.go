package main

import "fmt"

type Person struct {
	name string
	age int
	sex byte
}

type Student struct {
	Person
	id int
	addr string
	name string
}

func main() {
	var s Student
	s.name = "ksm"
	s.sex = 'g'
	s.age = 18
	s.addr = "suzhou"
	s.Person.name = "xiaomili"

	fmt.Printf("s = %+v\n", s)
}
