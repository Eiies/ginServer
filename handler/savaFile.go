package handler

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

// FolderPath 文件夹路径
const FolderPath = "./uploads"

func SaveFile(fileHeader *multipart.FileHeader) (string, string, error) {
	folder := FolderPath
	// 基于时间戳生成文件名
	fileName := strconv.FormatInt(time.Now().UnixNano(), 10)
	// 检查文件夹是否存在，不存在则创建
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			log.Fatalln("无法创建文件夹：", err)
		}
	}

	// 创建目标文件
	dst, err := os.Create(fmt.Sprintf("%s/%s", folder, fileName))
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

	// 将上传的文件内容复制到目标文件，并同时保存为Base64编码
	encoder := base64.NewEncoder(base64.StdEncoding, dst)
	defer encoder.Close()

	// 将上传的文件内容复制到Base64编码器中
	_, err = io.Copy(encoder, src)
	if err != nil {
		return "", "", err
	}

	return folder, fileName, nil
}
