package utils

type Operation int

const (
	Set Operation = iota + 1
	Get
	Del
)
