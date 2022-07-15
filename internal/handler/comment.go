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
		response.NewError(c, http.StatusBadRequest, "invalid post id param")
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
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(c, http.StatusBadRequest, "invalid post id param")
		return
	}

	comments, err := h.services.Comment.GetAll(userId, postId)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *Handler) getCommentById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(c, http.StatusBadRequest, "invalid comment id param")
		return
	}

	comment, err := h.services.Comment.GetById(userId, commentId)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (h *Handler) updateComment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.UpdateCommentInput
	if err := c.BindJSON(&input); err != nil {
		response.NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Comment.Update(userId, id, input); err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.StatusResponse{
		Status: http.StatusText(200),
	})
}

func (h *Handler) deleteComment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(c, http.StatusBadRequest, "invalid comment id param")
		return
	}

	err = h.services.Comment.Delete(userId, commentId)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.StatusResponse{
		Status: http.StatusText(200),
	})
}
