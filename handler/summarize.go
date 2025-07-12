package handler

import (
	"go-ai-summarizer/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SummarizeRequest struct {
	Text string `json:"text"`
}

func SummarizeHandler(c *gin.Context) {
	var req SummarizeRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	summary, err := service.CallDeepSeekAPI(req.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"summary": summary})
}
