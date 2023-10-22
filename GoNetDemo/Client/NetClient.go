package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddr := "127.0.0.1:8080"
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("连接服务端失败：", err)
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		// 读取数据
		fmt.Print("请输入消息：")
		msg, _ := reader.ReadString('\n')

		// 发送数据
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("发送数据出错：", err)
			return
		}

		// 接收数据
		respData := make([]byte, 1024)
		n, err := conn.Read(respData)
		if err != nil {
			fmt.Println("接收数据出错：", err)
			return
		}
		resp := string(respData[:n])

		// 处理粘包
		fmt.Println("接收到响应：", resp[:len(resp)-1])
	}
}
