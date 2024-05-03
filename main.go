package main

import (
	"fmt"
	"goServer/handler"
	"goServer/router"
	"log"
)

func main() {
	r := router.SetupRouter()
	// 输出本机 IP
	for _, ip := range handler.GetLocalIPv4() {
		fmt.Printf("\nhttp://%v:9090/api\n", ip)
	}
	// 启动服务，端口为 9090
	err := r.Run(fmt.Sprintf(":%d", 9090))
	if err != nil {
		log.Fatal(err)
	}
}
