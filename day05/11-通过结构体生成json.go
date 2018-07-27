package main

import (
	"fmt"
	"encoding/json"
)

type IT struct {
	Company string `json:"-"`
	Subjects []string `json:"subject"`
	IsOK bool `json:",string"`
	Price float64 `json:"price"`
}

func main() {
	it := IT{
		"itcast",
		[]string{"Go", "C++", "Python", "Golang"},
		true,
		666.666,
	}

	//slice, err := json.Marshal(it)
	//if err != nil {
	//	fmt.Println("err = ", err)
	//	return
	//}
	//fmt.Println("slice = ", string(slice))

	slice, err := json.MarshalIndent(it, "", "	 ")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("slice = ", string(slice))
}
