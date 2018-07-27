package main

import "fmt"

func test01()  {
	fmt.Println("111111")
}

func test02()  {
	fmt.Println("222222")
	panic("this is a panic test!")
}

func test03()  {
	fmt.Println("333333")
}

func main() {
	test01()
	test02()
	test03()
}
