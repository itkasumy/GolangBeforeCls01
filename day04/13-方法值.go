package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

func (p Person) SetInfoValue()  {
	fmt.Printf("SetInfoValue: %p, %v\n", &p, p)
}

func (p *Person) SetInfoPointer()  {
	fmt.Printf("SetInfoPointer: %p, %v\n", p, p)
}

func main() {
	p := Person{"ksm", 'm', 18}

	fmt.Printf("main: %p, %v\n", &p, p)
	p.SetInfoPointer()

	pFunc := p.SetInfoPointer
	pFunc()

	vFunc := p.SetInfoValue
	vFunc()
}
