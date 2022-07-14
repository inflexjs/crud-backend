package service

import (
	"github.com/inflexjs/crud-backend/internal/models"
	"github.com/inflexjs/crud-backend/internal/storage"
)

type PostService struct {
	store storage.Post
}

func NewPostService(store storage.Post) *PostService {
	return &PostService{store: store}
}

func (s *PostService) Create(userId int, post models.Post) (int, error) {
	return s.store.Create(userId, post)
}
