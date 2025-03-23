package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	startServer()
}

func startServer() {
	r := gin.Default()

	r.GET("/api/time", getCurrentTime)
	r.GET("/api/getkey", getKey)
	r.POST("/api/setkey", setKey)

	slog.Info(" --- Started")
	r.Run(":9003")
}
