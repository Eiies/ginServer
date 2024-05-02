package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadResponseNotFound struct {
	Message int
}

type UploadResponseServerError struct {
	Message int
}

type UploadResponseOk struct {
	Message   int
	FilePath  string
	FileName  string
	Base64Enc bool
}

// AccessControlHeader 请求头
const AccessControlHeader = "Access-Control-Allow-Origin"

func UploadFile(c *gin.Context) {
	c.Header(AccessControlHeader, "*")    // 允许跨域请求
	fileHeader, err := c.FormFile("file") // 获取上传的文件和文件头信息

	// 上传失败，并返回错误 json
	if err != nil {
		statusNotFound := UploadResponseNotFound{
			Message: http.StatusNotFound,
		}
		c.JSON(http.StatusNotFound, statusNotFound)
		return
	}

	// 传递文件头信息给 SaveFile 函数
	_, fileName, err := SaveFile(fileHeader)
	if err != nil {
		statusServerError := UploadResponseServerError{
			Message: http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, statusServerError)
		return
	}

	// 构建成功响应
	statusOk := UploadResponseOk{
		Message:   http.StatusOK,
		FilePath:  "uploads",
		FileName:  fileName,
		Base64Enc: true,
	}

	// 返回成功响应
	c.JSON(http.StatusOK, statusOk)
}
