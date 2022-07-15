package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/inflexjs/crud-backend/internal/models"
	"github.com/inflexjs/crud-backend/internal/response"
)

func (h *Handler) createComment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.Comment
	if err := c.BindJSON(&input); err != nil {
		response.NewError(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Comment.Create(userId, postId, input)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllComments(c *gin.Context) {

}

func (h *Handler) getCommentById(c *gin.Context) {

}

func (h *Handler) updateComment(c *gin.Context) {

}

func (h *Handler) deleteComment(c *gin.Context) {

}
