package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func SendResult(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, Response{
		Data:    data,
		Message: message,
	})
}

func SendError(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Message: message,
	})
}
