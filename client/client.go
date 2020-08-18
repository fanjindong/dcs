package client

import (
	"bufio"
	"github.com/fanjindong/dcs/codec"
	"github.com/fanjindong/dcs/utils"
	"net"
)

type Client interface {
	Get(string) (string, error)
	Set(string, string) error
	Del(string) error
}

type TcpClient struct {
	conn    net.Conn
	encoder codec.Encoder
	decoder codec.Decoder
}

func NewTcpClient() *TcpClient {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	return &TcpClient{conn: conn, encoder: codec.NewDefaultEncoder(), decoder: codec.NewDefaultDecoder()}
}

func (c *TcpClient) Command(op utils.Operation, kv ...string) (string, error) {
	//defer c.conn.Close()
	body, err := c.encoder.Request(op, kv...)
	if err != nil {
		return "", err
	}
	if _, err := c.conn.Write(body); err != nil {
		return "", err
	}
	reader := bufio.NewReader(c.conn)
	return c.decoder.Response(reader)
}
