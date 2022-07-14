package storage

import (
	"github.com/inflexjs/crud-backend/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
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

// docker run --name=crud-db -e POSTGRES_PASSWORD='7970' -p 5436:5432 -d --rm postgres
