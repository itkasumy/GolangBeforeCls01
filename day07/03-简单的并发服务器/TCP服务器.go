package main

import (
	"net"
	"fmt"
	"strings"
)

func HandlerConn(conn net.Conn)  {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "addr connect successful")



	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))

		if "exit" == string(buf[:n - 1]) {
			fmt.Println(addr, " exit")
			return
		}

		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	listner, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listner.Close()

	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			continue
		}
		
		go HandlerConn(conn)
	}
}
