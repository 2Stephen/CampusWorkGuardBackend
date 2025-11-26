package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {
	c := AppConfig.Mysql

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}

	DB = db
	log.Println("数据库连接成功")
	sqlBytes, err := os.ReadFile("config/schema.sql")
	if err != nil {
		panic(err)
	}
	if err := DB.Exec(string(sqlBytes)).Error; err != nil {
		log.Printf("执行数据库初始化脚本失败: %v", err)
	}
	log.Println("数据库初始化完成")

}
