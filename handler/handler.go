package handler

import (
	"fmt"
	"github.com/fanjindong/dcs/cache"
	"github.com/fanjindong/dcs/codec"
	"github.com/fanjindong/dcs/utils"
	"net"
)

type Handler interface {
	Process(conn net.Conn)
}

type DefaultHandler struct {
	cache   cache.Cache
	encoder codec.Encoder
	decoder codec.Decoder
}

func NewDefaultHandler() Handler {
	return &DefaultHandler{cache: cache.NewMemoryCache(), encoder: codec.NewDefaultEncoder(), decoder: codec.NewDefaultDecoder()}
}

func (h *DefaultHandler) Process(conn net.Conn) {
	defer conn.Close()
	var (
		value = "ok"
	)

	op, kv, err := h.decoder.Request(conn)
	fmt.Println("DefaultHandler.Process.Request:", op, kv, err)
	if err != nil {
		_, _ = conn.Write(h.encoder.Response(err, ""))
	}

	switch op {
	case utils.Set:
		err = h.cache.Set(kv.GetKey(), kv.GetValue())
	case utils.Get:
		value, err = h.cache.Get(kv.GetKey())
	case utils.Del:
		err = h.cache.Del(kv.GetKey())
	}
	fmt.Println("DefaultHandler.Process.Response:", value, err)
	if err != nil {
		_, _ = conn.Write(h.encoder.Response(err, ""))
	}
	_, _ = conn.Write(h.encoder.Response(nil, value))
}
