package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string `json:company`
	Subjects []string `json:subjects`
	IsOK bool `json:isOK`
	Price float64 `json:price`
}

type IT2 struct {
	Subjects []string `json:subjects`
}

func main() {
	str := `
		{
	        "company": "itcast",
        	"isOK": true,
        	"price": 666.666,
        	"subjects": [
                	"Go",
                	"C++",
                	"Python",
            	    "Golang"
        	]
		}
	`

	var it IT
	err := json.Unmarshal([]byte(str), &it)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("it = ", it)
	fmt.Printf("it = %+v\n", it)

	var it2 IT2
	err2 := json.Unmarshal([]byte(str), &it2)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}
	fmt.Println("it2 = ", it2)
	fmt.Printf("it2 = %+v", it2)
}
