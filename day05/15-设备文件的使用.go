package main

import (
	"os"
	"fmt"
)

func main() {
	//os.Stdout.Close()
	fmt.Println("Are you ok ?")
	os.Stdout.WriteString("Are you ok ?\n")

	//os.Stdin.Close()
	var a int
	fmt.Println("请输入一个数字: ")
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
