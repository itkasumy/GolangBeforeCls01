package main

import "fmt"

func main()  {
	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	const d = iota
	fmt.Printf("d = %d\n", d)

	const (
		e = iota
		f
		g
	)
	fmt.Printf("e = %d, f = %d, g = %d\n", e, f, g)

	const (
		h = iota
		i, j, k = iota, iota, iota
		l, m = iota, iota
	)
	fmt.Printf("h = %d, i = %d, j = %d, k = %d, l = %d, m = %d", h, i, j, k, l, m)
}
