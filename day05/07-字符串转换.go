package main

import (
	"strconv"
	"fmt"
)

func main() {
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abcgohello")

	fmt.Println("slice = ", string(slice))

	var str string
	str = strconv.FormatBool(false)
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	str = strconv.Itoa(666)

	fmt.Println("str = ", str)


	flag, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag = ", flag)
	} else {
		fmt.Println("err = ", err)
	}

	a, _ := strconv.Atoi("1")
	fmt.Println("a = ", a)
}
