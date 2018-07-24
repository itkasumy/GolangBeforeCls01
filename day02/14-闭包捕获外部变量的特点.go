package main

import "fmt"

func main() {
	a := 10
	str := "ksm"

	func() {
		a = 20
		str = "kasumy"
		fmt.Printf("inner: a = %d, str = %s\n", a, str)
	}()
	fmt.Printf("outer: a = %d, str = %s\n", a, str)
}
