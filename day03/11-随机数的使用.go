package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		//fmt.Println("rand = ", rand.Int())
		fmt.Println("rand = ", rand.Intn(100))
	}
}
