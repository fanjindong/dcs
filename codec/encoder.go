package codec

import (
	"fmt"
	"github.com/fanjindong/dcs"
	"github.com/fanjindong/dcs/utils"
)

type Encoder interface {
	Request(operation utils.Operation, kv ...string) ([]byte, error)
	Response(err error, value string) []byte
}

type DefaultEncoder struct{}

func NewDefaultEncoder() *DefaultEncoder {
	return &DefaultEncoder{}
}

func (DefaultEncoder) Request(op utils.Operation, kv ...string) ([]byte, error) {
	switch op {
	case utils.Set:
		k, v := kv[0], kv[1]
		return []byte(fmt.Sprintf("%d%d %s%d %s", op, len(k), k, len(v), v)), nil
	case utils.Get:
		k := kv[0]
		return []byte(fmt.Sprintf("%d%d %s", op, len(k), k)), nil
	case utils.Del:
		k := kv[0]
		return []byte(fmt.Sprintf("%d%d %s", op, len(k), k)), nil
	default:
		return nil, dcs.OpNotExistsErr
	}
}

func (e DefaultEncoder) Response(err error, value string) []byte {
	if err == nil {
		return []byte(fmt.Sprintf("0 %d %s", len(value), value))
	} else {
		return []byte(fmt.Sprintf("%d %s%d %s", len(err.Error()), err.Error(), len(value), value))
	}
}
