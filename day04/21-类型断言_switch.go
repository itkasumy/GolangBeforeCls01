package main

import "fmt"

type Student struct {
	id int
	name string
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "Golang"
	i[2] = Student{666, "ksm"}

	for index, data := range i {
		switch val := data.(type) {
		case int:
			fmt.Printf("i[%d]类型为int, 内容为%d\n", index, val)
		case string:
			fmt.Printf("i[%d]类型为string,内容为%s\n", index, val)
		case Student:
			fmt.Printf("i[%d]类型为Student,内容为name = %s, id = %d\n", index, val.name, val.id)
		}
	}
}
