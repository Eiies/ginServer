package router

import (
	handler "goServer/handler/file"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由函数
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(corsMiddleware())
	// api 主路由
	api := router.Group("/api")
	{
		// 主路由下的子路由 /upload
		api.POST("/upload", handler.UploadFile)
		api.GET("/readfile", handler.ReadFile)

	}
	// 给网页返回一个 favicon.ico
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.File("public/favicon.ico")
	})
	return router
}

// CORS 中间件
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // 缓存预检请求结果的时间，单位秒
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	}
}
