package main

import "fmt"

func main()  {
	var a int
	fmt.Println("请输入变量a:")

	//fmt.Scanf("%d", &a)
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
