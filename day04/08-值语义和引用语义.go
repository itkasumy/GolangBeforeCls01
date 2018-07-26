package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

func (p Person) SetInfoValue(n string, s byte, a int)  {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("SetInfoValue &p = %p\n", &p)
}

func (p *Person) SetInfoPointer(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("SetInfoPointer p = %p\n", p)
}

func main()  {
	p := Person{"go", 'n', 17}
	fmt.Printf("&p = %p\n", &p)

	p.SetInfoValue("ksm", 'm', 18)
	fmt.Printf("&p = %p", &p)
}
