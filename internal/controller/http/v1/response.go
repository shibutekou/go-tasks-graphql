package v1

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	slog.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
