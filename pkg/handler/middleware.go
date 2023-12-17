package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	apiKeyHeader        = "X-API-KEY"
)

func (h *Handler) userIdentity(c *gin.Context) {
	headerXkey := c.GetHeader(apiKeyHeader)
	headerBearer := c.GetHeader(authorizationHeader)

	if headerXkey != "" {
		userId, err := h.services.Authorization.GetUserIdByApiKey(headerXkey)
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		c.Set(userCtx, userId)

		return

	} else if headerBearer != "" {
		headerParts := strings.Split(headerBearer, " ")
		if len(headerParts) != 2 {
			newErrorResponse(c, http.StatusUnauthorized, "invalid auth headerBearer")
			return
		}

		userId, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		c.Set(userCtx, userId)

		return

	} else {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
