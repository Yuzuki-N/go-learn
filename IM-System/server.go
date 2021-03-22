package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip string
	Port int

	OnlineMap map[string]*User
	maplock sync.RWMutex

	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip: ip,
		Port: port,
		OnlineMap: make(map[string]*User),
		Message: make(chan string),
	}
	return server
}

//监听广播消息,有消息就发给全部的在线user
func(server *Server) ListenMessage() {
	for {
		msg := <- server.Message

		server.maplock.Lock()
		for _, cli := range server.OnlineMap {
			cli.C <- msg
		}
		server.maplock.Unlock()
	}
}


func(server *Server) Broadcast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ": " + msg

	server.Message <- sendMsg
}

func (server *Server) Handler(conn net.Conn) {
	user := NewUser(conn)

	server.maplock.Lock()
	server.OnlineMap[user.Name] = user
	server.maplock.Unlock()


	server.Broadcast(user, "已上线")

	select {

	}
}

func (server *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("net.Listener err...: ", err)
		return
	}
	defer listener.Close()

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