package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}
