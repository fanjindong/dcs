package main

import "github.com/fanjindong/dcs/service"

func main() {
	s := service.NewTcpService()
	s.Run()
}
