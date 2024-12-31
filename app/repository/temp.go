package repository

import (
	"time"
)

type TempMemoryRepo interface {
	Set(key string, value string, expiration time.Duration) error
	Get(key string) (string, error)
	List(skip, limit int) ([]string, error)
}

type tempRepositoryImpl struct {
	Temp TempMemoryRepo
}

func NewTempMemory(temp TempMemoryRepo) TempMemoryRepo {
	return &tempRepositoryImpl{Temp: temp}
}

func (r *tempRepositoryImpl) Set(key, value string, expiration time.Duration) error {
	return r.Temp.Set(key, value, expiration)
}

func (r *tempRepositoryImpl) Get(key string) (string, error) {
	return r.Temp.Get(key)
}

func (r *tempRepositoryImpl) List(skip, limit int) ([]string, error) {
	return r.Temp.List(skip, limit)
}
