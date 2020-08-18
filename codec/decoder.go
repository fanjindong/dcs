package codec

import (
	"bufio"
	"errors"
	"github.com/fanjindong/dcs/utils"
	"strconv"
)

type Decoder interface {
	Request(reader *bufio.Reader) (utils.Operation, *utils.Kv, error)
	Response(reader *bufio.Reader) (string, error)
}

type DefaultDecoder struct{}

func NewDefaultDecoder() *DefaultDecoder {
	return &DefaultDecoder{}
}

func (d DefaultDecoder) Request(reader *bufio.Reader) (utils.Operation, *utils.Kv, error) {
	var (
		op    utils.Operation
		key   string
		value string
		err   error
	)

	if op, err = d.readOp(reader); err != nil {
		return 0, nil, err
	}
	if key, err = d.readKey(reader); err != nil {
		return op, nil, err
	}
	switch op {
	case utils.Set:
		if value, err = d.readValue(reader); err != nil {
			return op, nil, err
		}
	}
	return op, utils.NewKv(key, value), nil
}

func (d DefaultDecoder) Response(reader *bufio.Reader) (string, error) {
	var (
		errMsg string
		value  string
		err    error
	)
	if errMsg, err = d.readError(reader); err != nil {
		return value, err
	}
	if value, err = d.readValue(reader); err != nil {
		return value, err
	}
	if errMsg == "" {
		return value, nil
	} else {
		return value, errors.New(errMsg)
	}
}

func (DefaultDecoder) readOp(r *bufio.Reader) (utils.Operation, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	op, err := utils.ByteToInt(b)
	if err != nil {
		return 0, err
	}
	return utils.Operation(op), nil
}

func (d DefaultDecoder) readKey(r *bufio.Reader) (string, error) {
	length, err := d.readLength(r)
	if err != nil {
		return "", err
	}
	key := make([]byte, length)
	if _, err := r.Read(key); err != nil {
		return "", err
	}
	return string(key), nil
}

func (d DefaultDecoder) readValue(r *bufio.Reader) (string, error) {
	length, err := d.readLength(r)
	if err != nil {
		return "", err
	}
	value := make([]byte, length)
	if _, err := r.Read(value); err != nil {
		return "", err
	}
	return string(value), nil
}

func (d DefaultDecoder) readError(r *bufio.Reader) (string, error) {
	length, err := d.readLength(r)
	if err != nil || length == 0 {
		return "", err
	}
	key := make([]byte, length)
	if _, err := r.Read(key); err != nil {
		return "", err
	}
	return string(key), nil
}

func (DefaultDecoder) readLength(r *bufio.Reader) (int, error) {
	line, err := r.ReadSlice(' ')
	if err != nil {
		return 0, err
	}
	if length, err := strconv.ParseInt(string(line[:len(line)-1]), 10, 64); err != nil {
		return 0, err
	} else {
		return int(length), nil
	}
}
