package handler

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func ReadFile(c *gin.Context) {
	c.Header(AccessControlHeader, "*") // 允许跨域请求
	// 存储文件名与文件内容的对应关系
	fileMap := make(map[string]string)

	// 循环查看文件夹里面的文件
	files, err := os.ReadDir(FolderPath)
	if err != nil {
		log.Fatal("查看文件名出错：", err)
	}
	for _, file := range files {
		if !file.IsDir() {
			// 构建完整的文件路径
			filePath := filepath.Join(FolderPath, file.Name())

			// 读取文件内容
			fileContent, err := os.ReadFile(filePath)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// 将文件名和文件内容存储到 map 中
			fileMap[file.Name()] = string(fileContent)
		}
	}
	// 将文件名和文件内容以 JSON 返回给前端
	c.JSON(http.StatusOK, fileMap)
}
