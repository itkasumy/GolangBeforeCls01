package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
	fmt.Printf("swap: a = %d, b = %d\n", *a, *b)
}

func main()  {
	a, b := 10, 20
	fmt.Printf("main before: a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("main after: a = %d, b = %d\n", a, b)
}
