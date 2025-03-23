package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	key   = "key"
	value = "123"
)

func sendInternalError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error: " + err.Error()})
}

func getCurrentTime(c *gin.Context) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	c.JSON(http.StatusOK, gin.H{"current_time": currentTime})
}

func getKey(c *gin.Context) {
	rdb, err := NewRedisClient()
	if err != nil {
		slog.Error("Error connecting to Redis", "error", err)
		sendInternalError(c, err)
		return
	}

	value, err := rdb.ReadValue(key)
	if err != nil {
		slog.Error("Error getting value from Redis")
		sendInternalError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
}

func setKey(c *gin.Context) {
	rdb, err := NewRedisClient()
	if err != nil {
		slog.Error("Error connecting to Redis", "error", err)
		sendInternalError(c, err)
		return
	}

	err = rdb.WriteValueWithTTL(key, value, time.Minute)
	if err != nil {
		slog.Error("Error getting value from Redis")
		sendInternalError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
}
