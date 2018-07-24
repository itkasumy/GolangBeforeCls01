package main

import "fmt"

func main()  {
	type bigint int64
	var a bigint
	fmt.Printf("type of a is %T\n", a)

	type (
		long int
		char byte
	)
	var b long = 11
	var ch char = 'a'
	fmt.Printf("b = %d, ch = %c\n", b, ch)
}
