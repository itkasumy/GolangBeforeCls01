package main

import (
	"encoding/json"
	"fmt"
)

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

	m := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//fmt.Println("m = ", m)
	//fmt.Printf("m = %+v\n", m)

	var res interface{}
	res = m["company"]
	fmt.Println("res = ", res)
	//for key, val := range m {
	//	fmt.Printf("%v => %v\n", key, val)
	//	//if key == "company" {
	//	//	res = val
	//	//	fmt.Println("res = ", res)
	//	//}
	//	switch data := val.(type) {
	//	case string:
	//		fmt.Printf("map[%s]的值类型为 string, val = %s\n", key, data)
	//	case bool:
	//		fmt.Printf("map[%s]的值类型为 bool, val = %v\n", key, data)
	//	case float64:
	//		fmt.Printf("map[%s]的值类型为 float64, val = %f\n", key, data)
	//	case []interface{}:
	//		fmt.Printf("map[%s]的值类型为 []interface, val = %v\n", key, data)
	//	}
	//}
}
