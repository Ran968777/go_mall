package main

import (
	"log"
	"strconv"

	"go-mall/config"
	"go-mall/routes"
)

func main() {
	// 加载配置
	cfg := config.GetConfig()

	// 初始化路由
	r := routes.SetupRouter()

	// 启动服务器
	port := strconv.Itoa(cfg.Server.Port)
	log.Printf("服务器启动在端口: %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
