package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	list := os.Args
	if len(list) != 3 {
		fmt.Println("usage: xxx srcFile dstFile")
		return
	}

	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName {
		fmt.Println("目的文件名不能和原文件名相同")
		return
	}

	sF, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	dF, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}

	defer sF.Close()
	defer dF.Close()

	buf := make([]byte, 4 * 1024)
	for {
		n, err := sF.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}
		dF.Write(buf[:n])
	}
}
