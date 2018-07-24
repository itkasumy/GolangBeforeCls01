package main

import "fmt"

func main()  {
	fmt.Printf("大写: %d, 小写: %d\n", 'A', 'a')
	fmt.Printf("大写转小写: %c\n", 'A' + 32)
	fmt.Printf("小写转大写: %c\n", 'a' - 32)
}
