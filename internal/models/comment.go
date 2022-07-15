package models

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
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
