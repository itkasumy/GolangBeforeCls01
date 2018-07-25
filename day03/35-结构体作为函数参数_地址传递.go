package main

import "fmt"

type Student struct {
	id int
	name string
	age int
	sex byte
	addr string
}

func Test(s *Student)  {
	s.id = 666
	fmt.Println("Test: s = ", s)
}

func main()  {
	s := Student{1, "ksm", 18, 'g', "wuhan"}
	fmt.Println("main before: s = ", s)

	Test(&s)
	fmt.Println("main after: s = ", s)
}
