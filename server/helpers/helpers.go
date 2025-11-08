package helpers

import (
	"github.com/gin-gonic/gin"
)

func InternalServerError(c *gin.Context) {
	c.JSON(500, gin.H{
		"success": false,
		"message": "Internal server error",
	})
}
