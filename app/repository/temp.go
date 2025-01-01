package repository

import (
	"time"
)

type TempMemoryRepository interface {
	Set(key string, value string, expiration time.Duration) error
	Get(key string) (string, error)
	List(skip, limit int) ([]map[string]interface{}, error)
	AllList() ([]map[string]interface{}, error)
}

type tempRepositoryImpl struct {
	Temp TempMemoryRepository
}

func NewTempMemory(temp TempMemoryRepository) TempMemoryRepository {
	return &tempRepositoryImpl{Temp: temp}
}

func (r *tempRepositoryImpl) Set(key, value string, expiration time.Duration) error {
	return r.Temp.Set(key, value, expiration)
}

func (r *tempRepositoryImpl) Get(key string) (string, error) {
	return r.Temp.Get(key)
}

func (r *tempRepositoryImpl) List(skip, limit int) ([]map[string]interface{}, error) {
	return r.Temp.List(skip, limit)
}

func (r *tempRepositoryImpl) AllList() ([]map[string]interface{}, error) {
	return r.Temp.AllList()
}
