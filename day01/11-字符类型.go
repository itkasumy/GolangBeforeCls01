package main

import "fmt"

func main()  {
	var ch byte
	ch = 97
	fmt.Println("ch = ", ch)

	fmt.Printf("%c, %d\n", ch, ch)

	ch = 'a'
	fmt.Printf("%c, %d\n", ch, ch)
}
