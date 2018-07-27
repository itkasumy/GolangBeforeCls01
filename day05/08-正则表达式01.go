package main

import (
	"regexp"
	"fmt"
)

func main() {
	str := "abc acc azc a7c 888 a9c tac"
	reg01 := regexp.MustCompile(`a.c`)
	if reg01 == nil {
		fmt.Println("regexp err")
		return
	}
	res := reg01.FindAllStringSubmatch(str, -1)
	fmt.Println("res = ", res)
}
