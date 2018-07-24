package main

import "fmt"

func fn(temp ...int)  {
	for i, data := range temp{
		fmt.Printf("temp[%d] = %d\n", i, data)
	}
}

func test(args ...int)  {
	fn(args...)
	fn(args[:2]...)
	fn(args[2:]...)
}

func main()  {
	test(1, 2, 3, 4)
}
