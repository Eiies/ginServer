package handler

import (
	"github.com/gin-gonic/gin"
)

type ErrorMessage struct {
	Message string `json:"Message"`
	Code    int    `json:"Code"`
	Error   string `json:"Error"`
}

// HandleError 错误处理
func HandleError(c *gin.Context, status int, err error, message string) {

	if err != nil {
		errorMessage := ErrorMessage{
			Message: message,
			Code:    status,
			Error:   err.Error(),
		}
		c.JSON(status, errorMessage)
		return
	}
}
