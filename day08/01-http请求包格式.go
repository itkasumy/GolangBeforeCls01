package main

import (
	"net"
	"fmt"
)

func main() {
	listenner, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	defer listenner.Close()

	conn, err01 := listenner.Accept()
	if err01 != nil {
		fmt.Println("listenner.Accept err01 = ", err01)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024 * 4)
	n, err02 := conn.Read(buf)
	if err02 != nil {
		fmt.Println("listenner.Accept err02 = ", err02)
		return
	}
	fmt.Println(string(buf[:n]))
}
