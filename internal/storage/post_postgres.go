package storage

import (
	"errors"
	"fmt"
	"strings"

	"github.com/inflexjs/crud-backend/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (r *PostPostgres) GetAll(userId int) ([]models.Post, error) {
	var posts []models.Post

	query := fmt.Sprintf(
		"SELECT tl.* FROM %s tl INNER JOIN %s ul on tl.id = ul.post_id WHERE ul.user_id = $1",
		postsTable, usersPostsTable,
	)

	err := r.db.Select(&posts, query, userId)

	return posts, err
}

func (r *PostPostgres) GetById(userId, postId int) (models.Post, error) {
	var post models.Post

	query := fmt.Sprintf(
		`SELECT tl.* FROM %s tl
				INNER JOIN %s ul on tl.id = ul.post_id WHERE ul.user_id = $1 AND ul.post_id = $2`,
		postsTable, usersPostsTable,
	)

	err := r.db.Get(&post, query, userId, postId)

	return post, err
}

func (r *PostPostgres) Delete(userId, postId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.post_id AND ul.user_id=$1 AND ul.post_id=$2",
		postsTable, usersPostsTable)

	result, err := r.db.Exec(query, userId, postId)

	row, _ := result.RowsAffected()
	if row == 0 {
		return errors.New("you cant delete this")
	}

	return err
}

func (r *PostPostgres) Update(userId, postId int, input models.UpdatePostInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Content != nil {
		setValues = append(setValues, fmt.Sprintf("content=$%d", argId))
		args = append(args, *input.Content)
		argId++
	}

	// title=$1
	// content=$1
	// title='uploaded title', content='uploaded content'
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.post_id AND ul.post_id=$%d AND ul.user_id=$%d",
		postsTable, setQuery, usersPostsTable, argId, argId+1)

	args = append(args, postId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)

	return err
}
