package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getCurrentTime(c *gin.Context) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	c.JSON(http.StatusOK, gin.H{"current_time": currentTime})
}

func main() {
	r := gin.Default()
	r.GET("/api/time", getCurrentTime)
	log.Println(" --- Started")
	r.Run(":9003")
}
