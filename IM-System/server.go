package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip string
	Port int

	OnlineMap map[string]*User
	mapLock   sync.RWMutex

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
func(server *Server) ListenMessager() {
	for {
		msg := <- server.Message

		server.mapLock.Lock()
		for _, cli := range server.OnlineMap {
			cli.C <- msg
		}
		server.mapLock.Unlock()
	}
}


func(server *Server) Broadcast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ": " + msg

	server.Message <- sendMsg
}

func (server *Server) Handler(conn net.Conn) {
	user := NewUser(conn)

	server.mapLock.Lock()
	server.OnlineMap[user.Name] = user
	server.mapLock.Unlock()


	server.Broadcast(user, "已上线")

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				server.Broadcast(user, "下线")
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err: ", err)
				return
			}
			//提取用户的消息，去除\n
			msg := string(buf[:n-1])
			server.Broadcast(user, msg)
		}
	}()

	select {}
}

func (server *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("net.Listener err...: ", err)
		return
	}
	defer listener.Close()

	go server.ListenMessager()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err: ", err)
			continue
		}
		go server.Handler(conn)

	}
}