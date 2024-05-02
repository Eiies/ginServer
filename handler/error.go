package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, status int, err error, message string) {
	if err != nil {
		log.Fatalln(err)
		c.JSON(status, gin.H{"error": message})
		return
	}
}
