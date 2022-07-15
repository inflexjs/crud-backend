package models

import "errors"

type Post struct {
	Id      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title" binding:"required"`
	Content string `json:"content" db:"content"`
}

type UsersPosts struct {
	Id     int
	UserId int
	PostId int
}

type UpdatePostInput struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

func (i UpdatePostInput) Validate() error {
	if i.Title == nil && i.Content == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
