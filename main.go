package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-ai-summarizer/config"
	"go-ai-summarizer/handler"
)

func main() {
	// 加载 .env 配置（API Key）
	config.LoadEnv()
	fmt.Println("当前API KEY:", config.GetDeepSeekKey())

	// 初始化 gin 路由
	r := gin.Default()

	// 定义 POST 接口 /summarize
	r.POST("/summarize", handler.SummarizeHandler)

	// 启动服务，监听 8080 端口
	r.Run(":8080")
}
