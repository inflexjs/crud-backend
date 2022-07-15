package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/inflexjs/crud-backend/internal/models"
	"github.com/inflexjs/crud-backend/internal/response"
)

func (h *Handler) createPost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input models.Post
	if err := c.BindJSON(&input); err != nil {
		response.NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Post.Create(userId, input)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllPosts(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	posts, err := h.services.Post.GetAll(userId)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.GetAllPostsResponse{
		Data: posts,
	})
}

func (h *Handler) getPostById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(c, http.StatusBadRequest, "invalid id param")
		return
	}

	post, err := h.services.Post.GetPostById(userId, id)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) updatePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.UpdatePostInput
	if err := c.BindJSON(&input); err != nil {
		response.NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(userId, id, input); err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.StatusResponse{
		Status: http.StatusText(200),
	})
}

func (h *Handler) deletePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Post.Delete(userId, id)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.StatusResponse{
		Status: http.StatusText(200),
	})
}
