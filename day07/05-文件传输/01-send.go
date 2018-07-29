package main

import (
	"fmt"
	"os"
	"net"
	"io"
)

func SendFile(path string, conn net.Conn)  {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024 * 4)

	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕！")
			} else {
				fmt.Println("f.Read err = ", err)
			}
			return
		}
		conn.Write(buf[:n])
	}
}

func main() {
	fmt.Println("请输入需要传输的文件:")
	var path string
	fmt.Scan(&path)

	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err = ", err)
		return
	}

	conn, err01 := net.Dial("tcp", "127.0.0.1:8000")
	if err01 != nil {
		fmt.Println("net.Dial err01 = ", err01)
		return
	}
	defer conn.Close()

	_, err02 := conn.Write([]byte(info.Name()))
	if err02 != nil {
		fmt.Println("conn.Write err02 = ", err02)
		return
	}

	buf := make([]byte, 1024)
	n, err03 := conn.Read(buf)
	if err03 != nil {
		fmt.Println("err03 = ", err03)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))
	if "ok" == string(buf[:n]) {
		SendFile(path, conn)
	}
}
