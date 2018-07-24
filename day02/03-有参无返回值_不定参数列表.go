package main

import "fmt"

func MyFn(args ...int) {
	fmt.Println("len(args) = ", len(args))

	for i := 0; i < len(args); i++ {
		fmt.Printf("arg[%d] = %d]\n", i, args[i])
	}

	fmt.Println("===============")

	for i, data := range args{
		fmt.Printf("arg[%d] = %d\n", i, data)
	}
}

func main()  {
	MyFn()
	MyFn(1)
	MyFn(1, 2)
	MyFn(1, 2, 3)
}
