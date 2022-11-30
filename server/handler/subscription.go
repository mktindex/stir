package handler

import "github.com/gin-gonic/gin"

func GetSubscriptionHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "获取订阅列表"})
}

func CreateSubscriptionHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "创建订阅"})
}
