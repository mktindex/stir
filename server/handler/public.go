package handler

import "github.com/gin-gonic/gin"

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func ClientIPHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": c.ClientIP()})
}

func NoticeHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "公告信息"})
}

func ServerStatusHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "服务器状态"})
}

func VersionHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": map[string]string{"version": "0.1.0"}})
}

func LatestMonthHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "返回数据可用日期时间"})
}

func ExchangeHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "汇率相关"})
}
