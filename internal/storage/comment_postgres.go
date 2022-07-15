package storage

import (
	"fmt"
	"strings"

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

	createCommentQuery := fmt.Sprintf("INSERT INTO %s (content) values ($1) RETURNING id", commentsTable)
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

func (r *CommentPostgres) GetAll(userId, postId int) ([]models.Comment, error) {
	var comments []models.Comment
	query := fmt.Sprintf("SELECT ti.* FROM %s ti INNER JOIN %s li on li.comment_id = ti.id INNER JOIN %s ul on ul.post_id = li.post_id WHERE li.post_id = $1 AND ul.user_id = $2",
		commentsTable, postsCommentsTable, usersPostsTable)

	if err := r.db.Select(&comments, query, postId, userId); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentPostgres) GetById(userId, commentId int) (models.Comment, error) {
	var comment models.Comment
	query := fmt.Sprintf("SELECT ti.* FROM %s ti INNER JOIN %s li on li.comment_id = ti.id INNER JOIN %s ul on ul.post_id = li.post_id WHERE ti.id = $1 AND ul.user_id = $2",
		commentsTable, postsCommentsTable, usersPostsTable)

	if err := r.db.Get(&comment, query, commentId, userId); err != nil {
		return comment, err
	}

	return comment, nil
}

// TODO rework all db strings (ti, li, ul)

func (r *CommentPostgres) Delete(userId, commentId int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s li, %s ul WHERE ti.id = li.comment_id AND li.post_id = ul.post_id AND ul.user_id = $1 AND ti.id = $2",
		commentsTable, postsCommentsTable, usersPostsTable)

	_, err := r.db.Exec(query, userId, commentId)

	return err
}

func (r *CommentPostgres) Update(userId, commentId int, input models.UpdateCommentInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Content != nil {
		setValues = append(setValues, fmt.Sprintf("content=$%d", argId))
		args = append(args, *input.Content)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id = li.comment_id AND li.post_id = ul.post_id AND ul.user_id=$%d AND ti.id = $%d",
		commentsTable, setQuery, postsCommentsTable, usersPostsTable, argId, argId+1)

	args = append(args, userId, commentId)

	_, err := r.db.Exec(query, args...)

	return err
}
