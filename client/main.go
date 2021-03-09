package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

const server = "localhost:8080"

//模拟客户端
func main() {

	addr, err := net.ResolveTCPAddr("tcp4", server)
	checkError(err)

	//建立tcp连接
	conn, err := net.DialTCP("tcp4", nil, addr)
	checkError(err)

	//向服务端发送数据
	_, err = conn.Write([]byte("hello !!!"))
	checkError(err)
	//接收响应
	response, _ := ioutil.ReadAll(conn)
	fmt.Println(string(response))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
