package storage

import (
	"errors"
)

type MemoryStorage struct {
	data map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{data: make(map[string]string)}
}

func (ms *MemoryStorage) Get(key string) (string, error) {
	value, ok := ms.data[key]

	if !ok {
		return value, errors.New("not found")
	}

	return value, nil
}

func (ms *MemoryStorage) Set(key string, value string) {
	ms.data[key] = value
}

func (ms *MemoryStorage) Delete(key string) error {
	_, ok := ms.data[key]

	if !ok {
		return errors.New("not found")
	}

	delete(ms.data, key)
	return nil
}
