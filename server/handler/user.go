package handler

import "github.com/gin-gonic/gin"

func TranslationHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "翻译相关"})
}
