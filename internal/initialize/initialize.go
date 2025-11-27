package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Mysql struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
	} `mapstructure:"mysql"`

	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
		PoolSize int    `yaml:"pool_size"`
	} `mapstructure:"redis"`

	Email struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `mapstructure:"email"`

	JWTConfig struct {
		Secret  string `yaml:"secret"`
		Expires int    `yaml:"expires"`
	} `mapstructure:"jwt"`
}

var AppConfig Config

func InitConfig() {
	v := viper.New()
	v.SetConfigName("config.back")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config") // 配置文件位置

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s", err))
	}

	if err := v.Unmarshal(&AppConfig); err != nil {
		panic(fmt.Errorf("解析配置失败: %s", err))
	}

	log.Println("配置加载成功")
}
