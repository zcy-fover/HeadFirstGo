package main

import "net"

type User struct {
	Name    string
	Addr    string
	channel chan string
	conn    net.Conn
}

func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:    userAddr,
		Addr:    userAddr,
		channel: make(chan string),
		conn:    conn,
	}
	//启动当前 user 的监听器，监听 channel 消息
	go user.ListenMessage()
	return user
}

// ListenMessage 监听是否有消息，有就发送给客户端
func (user *User) ListenMessage() {
	for {
		msg := <-user.channel
		user.conn.Write([]byte(msg + "\n"))
	}
}
