package db

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const imagesTXT = "./db/imagePath.txt"

func LogocalDatabase(imagesPath string) {
	// 将路径中的反斜杠替换为正斜杠
	imagesPath = filepath.ToSlash(imagesPath)

	file, err := os.OpenFile(imagesTXT, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = fmt.Fprintf(writer, "%v\n\n", imagesPath)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	writer.Flush()
}

func ReadPath() ([]string, error) {
	// 读取文件内容
	data, err := os.ReadFile(imagesTXT)
	if err != nil {
		log.Fatal("无法读取文件：", err)
		return nil, err
	}

	// 将文件内容转换为字符串
	content := string(data)

	// 按行分割字符串
	lines := strings.Split(content, "\n")

	return lines, nil
}
