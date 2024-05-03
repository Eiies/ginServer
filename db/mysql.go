package db

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Mysql(c *gin.Context, filePath string) {

	const (
		username     = ""
		password     = ""
		tcp          = "tcp"
		host         = ""
		port         = "3306"
		databaseName = "Example"
	)
	// "username:password@tcp(host:port)/database_name"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%s)/%s", username, password, tcp, host, port, databaseName))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法连接数据库"})
		return
	}
	defer db.Close()

	// 插入文件路径到数据库
	_, err = db.Exec("INSERT INTO images (path) VALUES (?)", filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法插入数据库"})
		return
	}
	// 返回成功的响应
	c.JSON(http.StatusOK, gin.H{"message": "路径已记录在数据库中"})

}
