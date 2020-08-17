package service

import (
	"fmt"
	"github.com/fanjindong/dcs/handler"
	"net"
)

type Service interface {
	Run()
}
type TcpService struct {
}

func NewTcpService() *TcpService {
	return &TcpService{}
}

func (t TcpService) Run() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("tcp listen for: 8080")
	h := handler.NewDefaultHandler()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Tcp Accept Error:", err.Error())
			continue
		}
		go h.Process(conn)
	}
}
