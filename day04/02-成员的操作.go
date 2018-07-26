package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

type Student struct {
	Person
	id int
	addr string
}

func main()  {
	var s Student = Student{Person{"ksm", 'g', 18}, 666, "wuhan"}
	s.name = "yoyo"
	s.sex = 'f'
	s.age = 24
	s.id = 777
	s.addr = "nanjing"

	s.Person = Person{"Java", 'n', 17}

	fmt.Println(s.name, s.sex, s.age, s.id, s.addr)
}
