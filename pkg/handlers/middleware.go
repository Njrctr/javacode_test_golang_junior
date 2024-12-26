package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	autorizationHeader = "Authorization"
	userCtx            = "userId"
	adminCtx           = "isAdmin"
)

func (h *Handler) userIdentify(c *gin.Context) {
	header := c.GetHeader(autorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid auth header")
		return
	}

	if headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid auth header")
		return
	}

	if headerParts[1] == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Token is empty")
		return
	}

	userId, isAdmin, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
	c.Set(adminCtx, isAdmin)
}

func (h *Handler) adminIdentify(c *gin.Context) {
	isAdmin, ok := c.Get(adminCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "Admin is not found")
		return
	}

	isAdminBool, ok := isAdmin.(bool)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "Admin is of invalid type")
		return
	}

	if !isAdminBool {
		newErrorResponse(c, http.StatusUnauthorized, "User is not admin")
		return
	}
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user Id is of invalid type")
	}

	return idInt, nil
}
