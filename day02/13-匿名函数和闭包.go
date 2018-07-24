package main

import "fmt"

func main() {
	a := 10
	str := "ksm"

	fn := func() {
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}

	fn()

	type fnType func()
	var fn2 fnType
	fn2 = fn
	fn2()

	func() {
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}()

	func(i, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}(1, 3)

	max, min := func(i, j int) (ma, mi int) {
		if i > j {
			ma = i
			mi = j
		} else {
			ma = j
			mi = i
		}
		return
	}(10, 20)
	fmt.Printf("max = %d, min = %d\n", max, min)
}
