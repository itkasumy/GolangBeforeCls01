package main

import (
	"fmt"
	"errors"
)

func main() {
	err1 := fmt.Errorf("%s", "this is a normal error01")
	fmt.Println("err1 = ", err1)

	err2 := errors.New("this is a normal error02")
	fmt.Println("err2 = ", err2)
}
