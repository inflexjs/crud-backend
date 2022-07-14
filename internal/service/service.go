package service

import "github.com/inflexjs/crud-backend/internal/storage"

type Authorization interface {
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
	return &Service{}
}
