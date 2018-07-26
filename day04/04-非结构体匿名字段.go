package main

import "fmt"

type MyStr string

type Person struct {
	name string
	age int
	sex byte
}

type Student struct {
	Person
	int
	MyStr
}

func main() {
	s := Student{Person{"ksm", 18, 'g'}, 666, "hahaha"}

	s.Person = Person{"xiaomili", 24, 'n'}

	fmt.Printf("s = %+v\n", s)
	fmt.Println(s.name, s.age, s.sex, s.int, s.MyStr)
	fmt.Println(s.Person, s.int, s.MyStr)
}
