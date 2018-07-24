package main

import "fmt"

func main()  {
	fmt.Println("111111")
	goto End
	fmt.Println("222222")
	End:
		fmt.Println("333333")
}
