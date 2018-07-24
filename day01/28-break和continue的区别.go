package main

import (
	"fmt"
	"time"
)

func main()  {
	i := 0
	for {
		i++
		time.Sleep(time.Second)
		if i == 5 {
			break
		}
		fmt.Println("i = ", i)
	}

	j := 0
	for {
		j++
		time.Sleep(time.Second)
		if j == 5 {
			continue
		}
		fmt.Println("j = ", j)
	}
}
