package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createPost(c *gin.Context) {
	id, _ := c.Get(userCtx)
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
