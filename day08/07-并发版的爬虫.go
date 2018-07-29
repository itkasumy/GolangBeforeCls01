package main

import (
	"fmt"
	"strconv"
	"net/http"
	"os"
)

func HttpGet(url string) (res string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	buf := make([]byte, 1024 * 4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.Read err = ", err)
			break
		}
		res += string(buf[:n])
	}
	return
}

func SpiderPage(i int, page chan<- int)  {
	//http://tieba.baidu.com/f/index/forumpark?cn=%E7%BE%8E%E5%89%A7&ci=0&pcn=%E7%94%B5%E8%A7%86%E5%89%A7&pci=0&ct=1&rn=20&pn=1
	//http://tieba.baidu.com/f/index/forumpark?cn=%E7%BE%8E%E5%89%A7&ci=0&pcn=%E7%94%B5%E8%A7%86%E5%89%A7&pci=0&ct=1&st=new&pn=2
	//http://tieba.baidu.com/f/index/forumpark?cn=%E7%BE%8E%E5%89%A7&ci=0&pcn=%E7%94%B5%E8%A7%86%E5%89%A7&pci=0&ct=1&st=new&pn=3
	//http://tieba.baidu.com/f/index/forumpark?cn=%E7%BE%8E%E5%89%A7&ci=0&pcn=%E7%94%B5%E8%A7%86%E5%89%A7&pci=0&ct=1&st=new&pn=4
	url := "http://tieba.baidu.com/f/index/forumpark?cn=%E7%BE%8E%E5%89%A7&ci=0&pcn=%E7%94%B5%E8%A7%86%E5%89%A7&pci=0&ct=1&st=new&pn=" + strconv.Itoa(i)
	//fmt.Println("url = ", url)
	res, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}

	fileName := strconv.Itoa(i) + ".html"
	f, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Println("os.Create err1 = ", err1)
		return
	}
	f.WriteString(res)
	defer f.Close()
	page<- i
}

func DoWork(start, end int)  {
	fmt.Printf("正在爬取 %d-%d 的页面\n", start, end)

	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}

func main() {
	var start, end int
	fmt.Println("请输入起始页（>=1）：")
	fmt.Scan(&start)
	fmt.Println("请输入终止页（>=起始页）：")
	fmt.Scan(&end)

	DoWork(start, end)
}
