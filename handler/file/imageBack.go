package handler

import (
	"fmt"
	"goServer/db"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReadFile 将图片返回给前端
func ReadImagesJSON(c *gin.Context) {

	imgs, _ := db.ReadPath()

	// 从图片列表中随机选择一张图片
	randomIndex := rand.Intn(len(imgs))
	randomImage := imgs[randomIndex]

	fmt.Println(imgs)

	// 构造 JSON 数据
	jsonData := gin.H{
		"Message": "NO",
		"image":   randomImage,
	}

	c.JSON(http.StatusOK, jsonData)
}

func ReadImages(c *gin.Context) {

	// 图片列表 imgs，包含多张图片的路径
	imgs, _ := db.ReadPath()

	// 从图片列表中随机选择一张图片
	randomIndex := rand.Intn(len(imgs))
	randomImage := imgs[randomIndex]

	// 返回图片
	c.File(randomImage)
}
