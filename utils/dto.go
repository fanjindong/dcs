package utils

type Kv struct {
	key   string
	value string
}

func NewKv(key string, value string) *Kv {
	return &Kv{key: key, value: value}
}

func (kv *Kv) GetKey() string {
	return kv.key
}

func (kv *Kv) GetValue() string {
	return kv.value
}
