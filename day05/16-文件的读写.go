package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

func WriteFile(path string)  {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()

	var str string
	for i := 0; i < 10; i++ {
		str = fmt.Sprintf("i = %d\n", i)
		n, err := f.WriteString(str)
		if err != nil {
			fmt.Println("err = ", err)
		}
		fmt.Println("n = ", n)
	}
}

func ReadFile(path string)  {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024 * 2)
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF {
		fmt.Println("err1 = ", err1)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))
}

func ReadFileLine(path string)  {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}
		fmt.Println("buf = ", string(buf))
	}
}

func main() {
	var path string = "./demo.txt"
	//WriteFile(path)
	
	//ReadFile(path)
	ReadFileLine(path)
}
