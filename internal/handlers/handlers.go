package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/skakunma/TestTaskTrood/internal/config"
)

func LoadHandlers(cfg *config.Config, c *gin.Engine) {
	c.POST("/send/message", CreateMessage(cfg))
}
