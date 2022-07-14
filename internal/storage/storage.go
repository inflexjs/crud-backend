package storage

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Storage{}
}

// docker run --name=crud-db -e POSTGRES_PASSWORD='7970' -p 5436:5432 -d --rm postgres
