package storage

import (
	"github.com/inflexjs/crud-backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Post interface {
	Create(userId int, post models.Post) (int, error)
	GetAll(userId int) ([]models.Post, error)
	GetPostById(userId, postId int) (models.Post, error)
	Delete(userId, postId int) error
	Update(userId, id int, input models.UpdatePostInput) error
}

type Comment interface {
	Create(postId int, comment models.Comment) (int, error)
}

type Storage struct {
	Authorization
	Post
	Comment
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Authorization: NewAuthPostgres(db),
		Post:          NewPostPostgres(db),
	}
}
