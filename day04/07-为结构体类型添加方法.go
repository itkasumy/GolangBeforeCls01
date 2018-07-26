package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

func (tmp Person) PrintInfo()  {
	fmt.Println("tmp = ", tmp)
}

func (p *Person) SetInfo(n string, s byte, a int)  {
	p.name = n
	p.sex = s
	p.age = a
}

func main() {
	p := Person{"ksm", 'm', 18}
	p.PrintInfo()

	var p2 Person
	(&p2).SetInfo("xiaomili", 'n', 18)
	p2.PrintInfo()
}
