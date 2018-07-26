package main

import "fmt"

type Person struct {
	name string
	age int
	sex byte
}

type Student struct {
	*Person
	id int
	addr string
}

func main() {
	s := Student{&Person{"xiaomili", 18, 'n'}, 666, "dongguan"}

	fmt.Println(s.name, s.age, s.sex, s.id, s.addr)

	var s1 Student
	s1.Person = new(Person)
	s1.name = "ksm"
	s1.age = 18
	s1.sex = 'g'
	fmt.Println(s1.name, s1.age, s1.sex)
}
