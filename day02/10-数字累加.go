package main

import "fmt"

func getSum() (sum int) {
	for i := 0; i <= 100; i++ {
		sum += i
	}
	return
}

func getSu(i int) (int) {
	if i == 1 {
		return i
	}
	return i + getSu(i - 1)
}

func main()  {
	sum := getSum()
	fmt.Println("sum = ", sum)

	su := getSu(100)
	fmt.Println("su = ", su)
}
