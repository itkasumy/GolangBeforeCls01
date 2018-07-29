package main

import (
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer conn.Close()

	conn.Write([]byte("Are you ok ?"))
}
