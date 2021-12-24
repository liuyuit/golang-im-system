package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 创建一个 Server
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}

	return server
}

func (this *Server) Handler(conn net.Conn) {
	// 当前连接的业务
	fmt.Println("连接建立成功")
}

// 前置要求，网络套接字编程
// 启动服务器
func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err ", err)
		return
	}
	defer listener.Close()
	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err: ", err)
			continue
		}

		// do handler
		go this.Handler(conn)
	}

	// close listen socket
}
