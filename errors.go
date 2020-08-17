package dcs

import "github.com/pkg/errors"

var (
	KeyNotExistsErr = errors.New("dcs: key not exists")
	OpNotExistsErr  = errors.New("dcs: operation not exists")
)
