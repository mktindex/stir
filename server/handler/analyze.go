package handler

import "github.com/gin-gonic/gin"

func CategoryAnalyzeHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "类目分析"})
}

func BrandAnalyzeHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "品牌分析"})
}
func AttributeAnalyzeHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "属性分析"})
}

func CustomAnalyzeHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "自定义分析"})
}
