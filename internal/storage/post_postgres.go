package storage

import (
	"fmt"

	"github.com/inflexjs/crud-backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (r *PostPostgres) Create(userId int, post models.Post) (int, error) {
	tx, err := r.db.Begin() // Транзакция
	if err != nil {
		return 0, err
	}

	var id int
	createPostQuery := fmt.Sprintf("INSERT INTO %s (title, content) VALUES ($1, $2) RETURNING id", postsTable)
	row := tx.QueryRow(createPostQuery, post.Title, post.Content)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersPostsQuery := fmt.Sprintf("INSERT INTO %s (user_id, post_id) VALUES ($1, $2)", usersPostsTable)
	_, err = tx.Exec(createUsersPostsQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
