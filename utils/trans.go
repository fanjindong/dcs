package utils

import (
	"strconv"
)

func ByteToInt(b byte) (int, error) {
	i64, err := strconv.ParseInt(string(b), 10, 64)
	return int(i64), err
}

func BytesToInt(bs []byte) (int, error) {
	i64, err := strconv.ParseInt(string(bs), 10, 64)
	return int(i64), err
}
