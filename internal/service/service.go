package service

import (
	"github.com/inflexjs/crud-backend/internal/models"
	"github.com/inflexjs/crud-backend/internal/storage"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type Post interface {
}

type Comment interface {
}

type Service struct {
	Authorization
	Post
	Comment
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Authorization: NewAuthService(storage.Authorization),
	}
}
