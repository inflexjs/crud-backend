package models

import "errors"

type Comment struct {
	Id      int    `json:"id" db:"id"`
	Content string `json:"content" db:"content" binding:"required"`
}

type UserComments struct {
	Id        int
	UserId    int
	CommentId int
}

type PostsComments struct {
	Id        int
	PostId    int
	CommentId int
}

type UpdateCommentInput struct {
	Content *string `json:"content"`
}

func (i UpdateCommentInput) Validate() error {
	if i.Content == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
