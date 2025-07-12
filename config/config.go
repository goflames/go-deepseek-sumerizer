package config

import (
	"log"
	"os"
)

func LoadEnv() {
	// 可选：你也可以检查变量是否存在
	if os.Getenv("DEEPSEEK_API_KEY") == "" {
		log.Println("警告：未检测到系统环境变量 DEEPSEEK_API_KEY")
	}
}

func GetDeepSeekKey() string {
	return os.Getenv("DEEPSEEK_API_KEY")
}
