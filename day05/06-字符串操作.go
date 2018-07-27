package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("hellomili", "hello"))
	fmt.Println(strings.Contains("hellomili", "xiaomili"))

	s := []string{"hello", "xiao", "mili"}
	str := strings.Join(s, " ")
	fmt.Println("str = ", str)

	fmt.Println(strings.Index(str, "mili"))
	fmt.Println(strings.Index(str, "xiaomili"))

	fmt.Println(strings.Repeat("junxia", 3))

	fmt.Println(strings.Split(str, " "))

	fmt.Println(strings.Trim(" Are you ok? ", " "))

	fmt.Println(strings.Fields(" Are you ok ? "))
}
