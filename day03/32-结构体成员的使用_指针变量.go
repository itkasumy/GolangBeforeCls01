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

	var p *Student
	p = &s
	p.id = 18
	(*p).name = "ksm"
	p.age = 24
	p.sex = 'g'
	p.addr = "beijing"

	fmt.Println("p = ", p)

	//通过new申请一个结构体
	s2 := new(Student)
	s2 = &s
	s2.id = 18
	s2.name = "ksm"
	s2.age = 24
	s2.sex = 'g'
	s2.addr = "beijing"

	fmt.Println("s2 = ", s2)

}
