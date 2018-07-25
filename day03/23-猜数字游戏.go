package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateRandNum(p *int)  {
	rand.Seed(time.Now().UnixNano())

	for {
		*p = rand.Intn(10000)
		if *p > 999 {
			break
		}
	}
}

func GetNum(s []int, num int)  {
	s[0] = num / 1000
	s[1] = num % 1000 / 100
	s[2] = num % 100 / 10
	s[3] = num % 10
}

func OnGame(randSlice []int)  {
	var num int
	keySlice := make([]int, 4)

	for {
		for {
			fmt.Println("请输入一个四位数: ")
			fmt.Scan(&num)
			if 999 < num && num < 10000 {
				break
			}
			fmt.Println("输入的数不符合要求！ ")
		}
		fmt.Println("num = ", num)

		GetNum(keySlice, num)
		fmt.Println("keySlice = ", keySlice)
		n := 0
		for i := 0; i < 4; i++ {
			if keySlice[i] > randSlice[i] {
				fmt.Printf("第 %d 位猜大了\n", i)
			} else if keySlice[i] < randSlice[i] {
				fmt.Printf("第 %d 位猜小了\n", i)
			} else {
				fmt.Printf("第 %d 位猜对了\n", i)
				n++
			}
		}

		if n == 4 {
			fmt.Println("全部猜对了")
			break
		}
	}
}

func main()  {
	var randNum int

	CreateRandNum(&randNum)

	//fmt.Println("randNum = ", randNum)

	randSlice := make([]int, 4)
	GetNum(randSlice, randNum)
	//fmt.Println("randSlice = ", randSlice)

	OnGame(randSlice)
}
