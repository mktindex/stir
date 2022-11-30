package handler

import "github.com/gin-gonic/gin"

func GetMarkHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "获取关注列表"})
}

func CreateMarkHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "创建关注"})
}
