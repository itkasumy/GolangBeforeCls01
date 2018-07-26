package main

import "fmt"

type Humaner interface {
	sayHi()
}

type Student struct {
	id int
	name string
}

type Teacher struct {
	addr string
	group string
}

type MyStr string

func (s *Student) sayHi() {
	fmt.Printf("Student: [%d, %s] sayhi\n", s.id, s.name)
}

func (t *Teacher) sayHi() {
	fmt.Printf("Teacher: [%s, %s] sayhi\n", t.addr, t.group)
}

func (m *MyStr) sayHi() {
	fmt.Printf("MyStr[%s] sayhi\n", *m)
}

func main() {
	var i Humaner

	s := &Student{666, "ksm"}
	i = s
	i.sayHi()

	t := &Teacher{"shenzhen", "FE"}
	i = t
	i.sayHi()

	var str MyStr = "hello"
	i = &str
	i.sayHi()

}
