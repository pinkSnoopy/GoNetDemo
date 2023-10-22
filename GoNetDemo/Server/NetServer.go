package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		// 读取数据
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取数据出错：", err)
			return
		}

		// 处理粘包
		msg := string(data[:len(data)-1])
		fmt.Println("接收到消息：", msg)

		// 发送数据
		resp := "已收到消息：" + msg + "\n"
		_, err = conn.Write([]byte(resp))
		if err != nil {
			fmt.Println("发送数据出错：", err)
			return
		}
	}
}

func main() {
	port := 8080
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("启动服务端失败：", err)
		return
	}

	fmt.Println("服务端已启动，监听端口：", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("建立连接出错：", err)
			continue
		}

		go handleConn(conn)
	}
}
