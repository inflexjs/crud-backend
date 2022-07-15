package service

import (
	"github.com/inflexjs/crud-backend/internal/models"
	"github.com/inflexjs/crud-backend/internal/storage"
)

type CommentService struct {
	commentStore storage.Comment
	postStore    storage.Post
}

func NewCommentService(commentStore storage.Comment, postStore storage.Post) *CommentService {
	return &CommentService{commentStore: commentStore, postStore: postStore}
}

func (s *CommentService) Create(userId, postId int, comment models.Comment) (int, error) {
	_, err := s.postStore.GetById(userId, postId)
	if err != nil {
		// list does not exist or does not belong to user
		return 0, err
	}

	return s.commentStore.Create(postId, comment)
}

func (s *CommentService) GetAll(userId, postId int) ([]models.Comment, error) {
	return s.commentStore.GetAll(userId, postId)
}

func (s *CommentService) GetById(userId, commentId int) (models.Comment, error) {
	return s.commentStore.GetById(userId, commentId)
}

func (s *CommentService) Delete(userId, commentId int) error {
	return s.commentStore.Delete(userId, commentId)
}

func (s *CommentService) Update(userId, commentId int, input models.UpdateCommentInput) error {
	return s.commentStore.Update(userId, commentId, input)
}
