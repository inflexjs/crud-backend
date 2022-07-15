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

func (s *PostService) GetAll(userId int) ([]models.Post, error) {
	return s.store.GetAll(userId)
}

func (s *PostService) GetById(userId, postId int) (models.Post, error) {
	return s.store.GetById(userId, postId)
}

func (s *PostService) Delete(userId, postId int) error {
	return s.store.Delete(userId, postId)
}

func (s *PostService) Update(userId, postId int, input models.UpdatePostInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.store.Update(userId, postId, input)
}
