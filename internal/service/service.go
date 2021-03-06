package service

import (
	"github.com/inflexjs/crud-backend/internal/models"
	"github.com/inflexjs/crud-backend/internal/storage"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Post interface {
	Create(userId int, post models.Post) (int, error)
	GetAll(userId int) ([]models.Post, error)
	GetById(userId, postId int) (models.Post, error)
	Delete(userId, postId int) error
	Update(userId, id int, input models.UpdatePostInput) error
}

type Comment interface {
	Create(userId, postId int, comment models.Comment) (int, error)
	GetAll(userId, postId int) ([]models.Comment, error)
	GetById(userId, commentId int) (models.Comment, error)
	Delete(userId, commentId int) error
	Update(userId, commentId int, input models.UpdateCommentInput) error
}

type Service struct {
	Authorization
	Post
	Comment
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Authorization: NewAuthService(storage.Authorization),
		Post:          NewPostService(storage.Post),
		Comment:       NewCommentService(storage.Comment, storage.Post),
	}
}
