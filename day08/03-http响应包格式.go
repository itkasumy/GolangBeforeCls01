package main

import (
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	defer conn.Close()

	requestBuf := ""
	conn.Write([]byte(requestBuf))

	buf := make([]byte, 1024 * 4)
	n, err01 := conn.Read(buf)
	if n == 0 {
		fmt.Println("Read err01 = ", err01)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))
}
