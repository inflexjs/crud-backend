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
	_, err := s.postStore.GetPostById(userId, postId)
	if err != nil {
		// list does not exist or does not belong to user
		return 0, err
	}

	return s.commentStore.Create(postId, comment)
}
