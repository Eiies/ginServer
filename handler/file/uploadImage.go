package handler

import (
	"goServer/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadImages 只能上传图片格式的文件
func UploadImages(c *gin.Context) {

	// 获取上传的文件
	file, err := c.FormFile("file")
	handler.HandleError(c, http.StatusBadRequest, err, "没有上传文件")
	SaveFile(file)

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"Message": "上传成功", "code": http.StatusOK})

}

func UploadFile(c *gin.Context) {

	// 获取上传的文件
	file, err := c.FormFile("file")
	handler.HandleError(c, http.StatusBadRequest, err, "没有上传文件")
	SaveFile(file)

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"Message": "上传成功", "code": http.StatusOK})

}
