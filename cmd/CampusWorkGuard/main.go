package main

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/router"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	initialize.InitConfig() // 读取 YAML
	initialize.InitDB()     // 初始化 GORM
	initialize.InitRedis()  // 初始化Redis连接配置
	err := godotenv.Load()  // 读取 .env 文件

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := router.SetupRouter()
	log.Fatal(r.Run(fmt.Sprintf(":%d", initialize.AppConfig.Server.Port)))
}
