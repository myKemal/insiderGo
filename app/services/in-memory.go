package services

import (
	"errors"
	"sync"
	"time"
)

type InMemoryClient struct {
	data  map[string]inMemoryValue
	mutex sync.RWMutex
}

type inMemoryValue struct {
	value      string
	expiration int64
}

func NewInMemoryClient() *InMemoryClient {
	return &InMemoryClient{
		data: make(map[string]inMemoryValue),
	}
}

func (c *InMemoryClient) Set(key string, value string, expiration time.Duration) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var exp int64
	if expiration > 0 {
		exp = time.Now().Add(expiration).Unix()
	}
	c.data[key] = inMemoryValue{
		value:      value,
		expiration: exp,
	}
	return nil
}

func (c *InMemoryClient) Get(key string) (string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	val, exists := c.data[key]
	if !exists {
		return "", errors.New("key not found")
	}

	if val.expiration > 0 && time.Now().Unix() > val.expiration {
		delete(c.data, key)
		return "", errors.New("key expired")
	}
	return val.value, nil
}

func (s *InMemoryClient) List(skip, limit int) ([]string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	keys := make([]string, 0, len(s.data))
	for key := range s.data {
		keys = append(keys, key)
	}

	if skip > len(keys) {
		return []string{}, nil
	}
	if skip+limit > len(keys) {
		limit = len(keys) - skip
	}

	return keys[skip : skip+limit], nil
}
