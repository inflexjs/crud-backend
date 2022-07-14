package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/inflexjs/crud-backend/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		posts := api.Group("/posts")
		{
			posts.POST("/", h.createPost)
			posts.GET("/", h.getAllPosts)
			posts.GET("/:id", h.getPostById)
			posts.PUT("/:id", h.updatePost)
			posts.DELETE("/:id", h.deletePost)

			comments := posts.Group(":id/comments")
			{
				comments.POST("/", h.createComment)
				comments.GET("/", h.getAllComments)
				comments.GET("/:comment_id", h.getCommentById)
				comments.PUT("/:comment_id", h.updateComment)
				comments.DELETE("/:comment_id", h.deleteComment)
			}
		}
	}

	return router
}
