package main

import "fmt"

func main()  {
	var num int
	fmt.Println("请输入楼层:")
	fmt.Scan(&num)
	switch num {
	case 1, 2, 3, 4:
		fmt.Printf("按下的是%d楼\n", num)
		fallthrough
	case 5, 6, 7, 8:
		fmt.Printf("按下的是%d楼\n", num)
	case 9, 10, 11, 12:
		fmt.Printf("按下的是%d楼\n", num)
	case 13, 14, 15, 16:
		fmt.Printf("按下的是%d楼\n", num)
	default:
		fmt.Println("按下的是0楼")
	}
}
