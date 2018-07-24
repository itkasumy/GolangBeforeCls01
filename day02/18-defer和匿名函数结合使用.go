package main

import "fmt"

func main()  {
	a, b := 10, 20
	defer func() {
		fmt.Printf("inner: a = %d, b = %d\n", a, b)
	}()


	defer func(a, b int) {
		fmt.Printf("inner width args: a = %d, b = %d\n", a, b)
	}(a, b)

	a, b = 111, 666
	fmt.Printf("outer: a = %d, b = %d\n", a, b)
}
