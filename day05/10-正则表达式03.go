package main

import (
	"regexp"
	"fmt"
)

func main() {
	str := `
		<head>
			<title>Go语言标准库文档中文版 | Go语言中文网 | Golang中文社区 | Golang中国</title>
			<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
			<meta http-equiv="X-UA-Compatible" content="IE=edge, chrome=1">
			<meta charset="utf-8">
			<link rel="shortcut icon" href="/static/img/go.ico">
			<link rel="apple-touch-icon" type="image/png" href="/static/img/logo2.png">
			<meta name="author" content="polaris <polaris@studygolang.com>">
			<meta name="keywords" content="中文, 文档, 标准库, Go语言,Golang,Go社区,Go中文社区,Golang中文社区,Go语言社区,Go语言学习,学习Go语言,Go语言学习园地,Golang 中国,Golang中国,Golang China, Go语言论坛, Go语言中文网">
			<meta name="description" content="Go语言文档中文版，Go语言中文网，中国 Golang 社区，Go语言学习园地，致力于构建完善的 Golang 中文社区，Go语言爱好者的学习家园。分享 Go 语言知识，交流使用经验">
		</head>
		<body>
			<div>123456</div>
			<div>一二三四五六</div>
			<div>
				Python
			</div>
			<div>热土</div>
			<div>山河</div>
		</body>
	`

	reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	if reg == nil {
		fmt.Println("MustCompile")
		return
	}

	res := reg.FindAllStringSubmatch(str, -1)
	fmt.Println("res = ", res)

	for _, text := range res {
		//fmt.Println("text[0] = ", text[0])
		fmt.Println("text[1] = ", text[1])
	}
}