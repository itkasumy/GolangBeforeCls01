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

func (p *Person) PrintInfo() {
	fmt.Printf("name=%s, sex=%c, age=%d\n", p.name, p.sex, p.age)
}

func (p *Student) PrintInfo() {
	fmt.Println("Student: p = ", p)
}

func main()  {
	s := Student{Person{"ksm", 'g', 18}, 666, "suzhou"}

	s.PrintInfo()

	s.Person.PrintInfo()
}
