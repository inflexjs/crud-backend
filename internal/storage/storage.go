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
}

type Comment interface {
}

type Storage struct {
	Authorization
	Post
	Comment
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Authorization: NewAuthPostgres(db),
	}
}
