/*整体概要*/

package handler

import "github.com/gin-gonic/gin"

func SummaryHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "整体概要"})
}

func SummaryTrendHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "整体趋势"})
}
