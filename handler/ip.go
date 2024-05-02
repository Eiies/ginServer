package handler

import (
	"log"
	"net"
)

// GetLocalIPv4 获取本机 IP
func GetLocalIPv4() []string {
	var ipv4Addrs []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipAddr := ipnet.IP.String()
				ipv4Addrs = append(ipv4Addrs, ipAddr)
			}
		}
	}
	if len(ipv4Addrs) == 0 {
		ipv4Addrs = append(ipv4Addrs, "获取 Ip 失败")
		return ipv4Addrs
	}
	// 返回切片数据
	return ipv4Addrs
}
