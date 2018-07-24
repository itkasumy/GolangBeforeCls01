package main

import "fmt"

func fnb() (b int) {
	b = 222
	fmt.Println("fnb b = ", b)
	return
}

func fna() (a int) {
	a = 111
	b := fnb()

	fmt.Println("fna b = ", b)

	fmt.Println("fna a = ", a)
	return
}

func main()  {
	fmt.Println("main func")
	a := fna()
	fmt.Println("main a = ", a)
}
