package handler

import (
	"go-ai-summarizer/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type URLRequest struct {
	URL string `json:"url"`
}

func UrlSummarizeHandler(c *gin.Context) {
	var req URLRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	// 提取正文并摘要
	summary, err := service.SummarizeURL(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"summary": summary})
}
