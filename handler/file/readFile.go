package handler

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func ReadFile(c *gin.Context) {
	files, err := os.ReadDir(FolderPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法从目录中读取文件"})
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(FolderPath, file.Name())
			c.File(filePath) // 将文件发送给前端
		}
	}

}
