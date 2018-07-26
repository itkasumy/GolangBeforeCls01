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

func WhoSayHi(i Humaner)  {
	i.sayHi()
}

func main() {
	s := &Student{666, "ksm"}
	t := &Teacher{"shenzhen", "FE"}
	var str MyStr = "hello"
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str
	for _, i := range x {
		i.sayHi()
	}
}
