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
	s1 := Student{1, "ksm", 18, 'g', "guangzhou"}
	s2 := Student{1, "ksm", 18, 'g', "guangzhou"}
	s3 := Student{1, "ksm", 18, 'g', "hangzhou"}
	var s4 Student = s3
	fmt.Println("s1 == s2", s1 == s2)
	fmt.Println("s1 == s3", s1 == s3)
	fmt.Println("s4 = ", s4)
}
