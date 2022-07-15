package storage

import (
	"fmt"

	"github.com/inflexjs/crud-backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type CommentPostgres struct {
	db *sqlx.DB
}

func NewCommentPostgres(db *sqlx.DB) *CommentPostgres {
	return &CommentPostgres{db: db}
}

func (r *CommentPostgres) Create(postId int, comment models.Comment) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var commentId int

	createCommentQuery := fmt.Sprintf("INSERT INTO %s (comment) values ($1) RETURNING id", commentsTable)
	row := tx.QueryRow(createCommentQuery, comment.Content)

	err = row.Scan(&commentId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createPostsCommentsQuery := fmt.Sprintf("INSERT INTO %s (post_id, comment_id) values ($1,$2)", postsCommentsTable)
	_, err = tx.Exec(createPostsCommentsQuery, postId, commentId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return commentId, tx.Commit()
}
