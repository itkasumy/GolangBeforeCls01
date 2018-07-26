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
		if val, ok := data.(int); ok {
			fmt.Printf("i[%d]类型为int, 内容为%d\n", index, val)
		} else if val, ok := data.(string); ok {
			fmt.Printf("i[%d]类型为string,内容为%s\n", index, val)
		} else if val, ok := data.(Student); ok {
			fmt.Printf("i[%d]类型为Student,内容为name = %s, id = %d\n", index, val.name, val.id)
		}
	}
}
