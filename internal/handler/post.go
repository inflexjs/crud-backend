package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inflexjs/crud-backend/internal/exception"
	"github.com/inflexjs/crud-backend/internal/models"
)

func (h *Handler) createPost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input models.Post
	if err := c.BindJSON(&input); err != nil {
		exception.NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Post.Create(userId, input)
	if err != nil {
		exception.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllPosts(c *gin.Context) {

}

func (h *Handler) getPostById(c *gin.Context) {

}

func (h *Handler) updatePost(c *gin.Context) {

}

func (h *Handler) deletePost(c *gin.Context) {

}
