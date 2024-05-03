package handler

import (
	"fmt"
	"goServer/db"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// FolderPath 文件夹路径
const FolderPath = "./uploads"

// SaveFile 保存文件
func SaveFile(fileHeader *multipart.FileHeader) (string, string, error) {

	// 获取上传文件的原始文件名和后缀
	fileExt := filepath.Ext(fileHeader.Filename)
	timeName := time.Now().UnixNano()

	// 基于时间戳生成文件名
	FileName := fmt.Sprintf("%d%s", timeName, fileExt)

	// 检查文件夹是否存在，不存在则创建
	if _, err := os.Stat(FolderPath); os.IsNotExist(err) {
		err := os.MkdirAll(FolderPath, 0755)
		if err != nil {
			log.Fatalln("无法创建文件夹：", err)
		}
	}

	// 创建目标文件
	dst, err := os.Create(fmt.Sprintf("%s/%s", FolderPath, FileName))
	if err != nil {
		log.Fatalln("创建文件失败", err)
	}

	defer dst.Close()

	// 打开上传的文件
	src, err := fileHeader.Open()
	if err != nil {
		return "", "", err
	}
	defer src.Close()

	/* 将上传的文件内容复制到目标文件，并同时保存为Base64编码
	encoder := base64.NewEncoder(base64.StdEncoding, dst)
	defer encoder.Close()
	*/

	// 将上传的文件内容复制
	_, err = io.Copy(dst, src)
	if err != nil {
		return "", "", err
	}

	db.LogocalDatabase(filepath.Join(dst.Name())) // 将文件路径储存到本地
	return FolderPath, FileName, nil
}
