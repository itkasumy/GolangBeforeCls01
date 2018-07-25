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
	var s1 *Student = &Student{1, "ksm", 18, 'n', "shenzhen"}
	fmt.Println("*s1 = ", *s1)

	s2 := &Student{name: "dongfangbubai", age: 18}
	fmt.Printf("type of s2 is %T\n", s2)
	fmt.Println("s2 = ", s2)
}
