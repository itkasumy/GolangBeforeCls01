package main

import (
	"net/http"
	"fmt"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http.Get err = ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status = ", resp.Status)
	fmt.Println("StatusCode = ", resp.StatusCode)
	fmt.Println("Header = ", resp.Header)
	fmt.Println("Body = ", resp.Body)

	var str string
	buf := make([]byte, 1024 * 4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.Read err = ", err)
			break
		}
		str += string(buf[:n])
	}
	fmt.Println("str = ", str)
}
