package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func RecvFile(fileName string, conn net.Conn)  {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024 * 4)

	for {
		n, err01 := conn.Read(buf)
		if err01 != nil {
			if err01 == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err01 = ", err01)
			}
			return
		}
		if n == 0 {
			fmt.Println("n == 0 文件接收完毕!")
			break
		}
		f.Write(buf[:n])
	}
}

func main() {
	listenner, err := net.Listen("tcp", "127.0.0.1:8000")
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

	buf := make([]byte, 1024)
	n, err02 := conn.Read(buf)
	if err02 != nil {
		fmt.Println("conn.Read err02 = ", err02)
		return
	}

	fileName := string(buf[:n])

	conn.Write([]byte("ok"))

	RecvFile(fileName, conn)
}
