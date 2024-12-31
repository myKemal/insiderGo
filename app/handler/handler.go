package handler

import (
	"github.com/myKemal/mongoApi/app/repository"
)

type Handler struct {
	Mongo repository.MongoRepository
	Temp  repository.TempMemoryRepo
}

func NewHandler(m repository.MongoRepository, t repository.TempMemoryRepo) *Handler {
	return &Handler{Mongo: m, Temp: t}
}
