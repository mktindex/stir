package handler

import "github.com/gin-gonic/gin"

func SettingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "配置相关"})
}
