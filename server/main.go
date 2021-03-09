package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

const port = "8080"

//模拟server端
func main() {
	tcpServer, _ := net.ResolveTCPAddr("tcp4", ":"+port)
	listener, _ := net.ListenTCP("tcp", tcpServer)

	for {
		//当有新的客户端请求来的时候，拿到与客户端的连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		//处理逻辑
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	//读取客户端传送的消息
	go func() {
		response, _ := ioutil.ReadAll(conn)
		fmt.Println(string(response))
	}()

	//向客户端发送消息
	time.Sleep(1 * time.Second)

	conn.Write([]byte("hello too"))
}
