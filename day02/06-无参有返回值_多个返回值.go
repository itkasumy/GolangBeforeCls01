package main

import "fmt"

func fn01() (int, int, int) {
	return 1, 3, 5
}

func fn02() (a int, b int, c int) {
	a = 111111
	b = 222222
	c = 333333
	return
}

func main()  {
	a, b, c := fn01()
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	c, d, e := fn02()
	fmt.Printf("c = %d, d = %d, e = %d\n", c, d, e)
}
