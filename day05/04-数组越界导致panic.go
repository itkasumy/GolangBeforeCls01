package main

import "fmt"

func test01()  {
	fmt.Println("111111")
}

func test02(x int)  {
	var arr [10]int
	arr[x] = 222
}

func test03()  {
	fmt.Println("333333")
}

func main() {
	test01()
	test02(10)
	test03()
}
