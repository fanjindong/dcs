package cache

import "github.com/fanjindong/dcs"

type Cache interface {
	Get(string) (string, error)
	Set(string, string) error
	Del(string) error
}

type MemoryCache struct {
	kv map[string]string
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{kv: map[string]string{}}
}

func (m *MemoryCache) Get(k string) (string, error) {
	if v, ok := m.kv[k]; ok {
		return v, nil
	} else {
		return "", dcs.KeyNotExistsErr
	}
}

func (m *MemoryCache) Set(k string, v string) error {
	m.kv[k] = v
	return nil
}

func (m *MemoryCache) Del(k string) error {
	delete(m.kv, k)
	return nil
}
