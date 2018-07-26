package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

func (p Person) SetInfoValue() {
	fmt.Println("SetInfoValue")
}

func (p *Person) SetInfoPointer() {
	fmt.Println("SetInfoPointer")
}

func main()  {
	p := Person{"ksm", 'g', 18}

	p.SetInfoPointer()
	p.SetInfoValue()
}
