package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skakunma/TestTaskTrood/internal/config"
	"github.com/skakunma/TestTaskTrood/internal/handlers"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		panic(err)
	}
	server := gin.Default()

	handlers.LoadHandlers(cfg, server)

	server.Run(cfg.AddrRunServer)
}
