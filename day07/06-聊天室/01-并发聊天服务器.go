package main

import (
	"net"
	"fmt"
	"strings"
	"time"
)

type Client struct {
	C chan string
	Name string
	Addr string
}

var onlineMap map[string]Client

var message = make(chan string)

func WriteMsgToClient(cli Client, conn net.Conn)  {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ": " + msg
	return
}

func HandleConn(conn net.Conn)  {
	defer conn.Close()

	cliAddr := conn.RemoteAddr().String()
	cli := Client{
		make(chan string),
		cliAddr,
		cliAddr,
	}

	onlineMap[cliAddr] = cli

	go WriteMsgToClient(cli, conn)

	//message <- "[" + cli.Addr + "]" + cli.Name + ": login"

	MakeMsg(cli, "login")
	cli.C<- MakeMsg(cli, "I am here")

	isQuit := make(chan bool)
	hasData := make(chan bool)

	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit<- true
				fmt.Println("conn.Read err = ", err)
				return
			}

			msg := string(buf[n - 1])

			if len(msg) == 3 && msg == "who" {
				conn.Write([]byte("user list: \n"))
				for _, tmp := range onlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 3 && msg[:6] == "rename" {
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("user list: \n"))
			}else {
				message<- MakeMsg(cli, msg)
			}
			hasData<- true
		}
	}()

	for {
		select {
		case isQuit:
			delete(onlineMap, cliAddr)
			message<- MakeMsg(cli, "login out")
			return
		case <-hasData:
		case <-time.After(time.Second * 60):
			delete(onlineMap, cliAddr)
			message<- MakeMsg(cli, "time out leave out")
			return
		}
	}
}

func Massager()  {
	onlineMap = make(map[string]Client)
	for {
		msg := <-message
		for _, cli := range onlineMap {
			cli.C <-msg
		}
	}
}

func main() {
	listenner, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	defer listenner.Close()
	
	go Massager()

	for {
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println("listenner.Accept err = ", err)
			continue
		}

		go HandleConn(conn)
	}
}
