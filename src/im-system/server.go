package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int
	//在线用户列表
	OnlineMap map[string]*User
	//锁
	mapLock sync.RWMutex
	//消息广播
	Message chan string
}

// NewServer 创建一个服务接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// Handler 处理当前业务
func (server *Server) Handler(conn net.Conn) {
	//创建一个用户
	user := NewUser(conn)
	//将用户放到在线列表中
	server.mapLock.Lock()
	server.OnlineMap[user.Name] = user
	server.mapLock.Unlock()

	//广播当前用户的上线消息
	server.BroadCast(user, "已上线")
	//将当前 go 程阻塞
	select {}
}

func (server *Server) BroadCast(user *User, message string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ": " + message
	//将广播消息放到消息 channel 中
	server.Message <- sendMsg
}

// ListenMessage 监听广播 message channel 一有消息 ，就发送给所有用户
func (server *Server) ListenMessage() {
	for {
		msg := <-server.Message
		server.mapLock.Lock()
		for _, cli := range server.OnlineMap {
			cli.channel <- msg
		}
		server.mapLock.Unlock()
	}
}

// Start 启动服务器接口
func (server *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("net.Listen err：", err)
		return
	}
	//退出时关闭监听
	defer listener.Close()
	//启动监听  message channel 的 go  程
	go server.ListenMessage()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err: ", err)
			continue
		}
		go server.Handler(conn)
	}
}
