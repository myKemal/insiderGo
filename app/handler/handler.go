package handler

import (
	"github.com/myKemal/insiderGo/app/repository"
)

type Handler struct {
	Mongo repository.MongoRepository
	Temp  repository.TempMemoryRepository
}

func NewHandler(m repository.MongoRepository, t repository.TempMemoryRepository) *Handler {
	return &Handler{Mongo: m, Temp: t}
}
