package main

import "fmt"

type Humaner interface {
	sayHi()
}

type Personer interface {
	Humaner
	sing(lrc string)
}

type Student struct {
	id int
	name string
}

func (s *Student) sayHi() {
	fmt.Printf("Student: [s.id, s.name] say hi\n", s.id, s.name)
}

func (s *Student) sing(lrc string) {
	fmt.Println("Student is singing ", lrc)
}

func main() {
	var iPro Personer
	iPro = &Student{666, "ksm"}
	var i Humaner
	i = iPro
	i.sayHi()
}
