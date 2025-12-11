package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
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

	// ========= 读取 schema.sql =========
	sqlBytes, err := os.ReadFile("config/schema.sql")
	if err != nil {
		panic(err)
	}

	sqlContent := strings.ReplaceAll(string(sqlBytes), "\r\n", "\n")

	// ========= 自动切分 SQL 语句 =========
	statements := splitSQL(sqlContent)

	// ========= 逐条执行，避免 Error 1064 =========
	for _, stmt := range statements {
		s := strings.TrimSpace(stmt)
		if s == "" {
			continue
		}
		log.Printf("执行 SQL:\n%s\n", s)

		// 单语句执行
		if err := DB.Exec(s).Error; err != nil {
			log.Printf("执行失败: %v\n", err)
		}
	}

	log.Println("数据库初始化完成")
}

/*
SQL 自动切分器
防止简单 split(";") 导致的函数体、触发器字符串等误切
*/
func splitSQL(sql string) []string {
	var stmts []string
	var buf strings.Builder
	inString := false

	for i := 0; i < len(sql); i++ {
		ch := sql[i]

		if ch == '\'' {
			inString = !inString
		}

		if ch == ';' && !inString {
			stmts = append(stmts, buf.String())
			buf.Reset()
			continue
		}

		buf.WriteByte(ch)
	}

	return stmts
}
