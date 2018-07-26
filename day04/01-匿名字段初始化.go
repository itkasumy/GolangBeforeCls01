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
	var s1 Student = Student{Person{"ksm", 'g', 18}, 666, "wuhan"}

	fmt.Println("s1 = ", s1)

	s2 := Student{Person{"mili", 'm', 18}, 777, "dongguan"}
	//fmt.Println("s2 = ", s2)
	fmt.Printf("s2 = %+v\n", s2)

	s3 := Student{id: 888}
	fmt.Printf("s3 = %+v\n", s3)

	s4 := Student{Person: Person{name: "mili"}, id: 1111}
	fmt.Printf("s4 = %+v", s4)
}
