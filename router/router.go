package router

import (
	"goServer/handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由函数
func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 添加路由// 主路由
	api := router.Group("/api")
	{
		// 主路由下的子路由 /upload
		api.POST("/upload", handler.UploadFile)
		api.GET("/readfile", handler.ReadFile)

	}

	return router
}
