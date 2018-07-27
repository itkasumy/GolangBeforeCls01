package main

import (
	"regexp"
	"fmt"
)

func main() {
	str := "3.14 567 asd 1.23 68.99 lkk 6.66 6.7 "
	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil {
		fmt.Println("MustCompile")
		return
	}

	res := reg.FindAllStringSubmatch(str, -1)
	fmt.Println("res = ", res)
}
