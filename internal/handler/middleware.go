package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/inflexjs/crud-backend/internal/exception"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentify(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		exception.NewError(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		exception.NewError(c, http.StatusUnauthorized, "invalid auth header")
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		exception.NewError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}
