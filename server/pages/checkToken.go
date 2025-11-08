package pages

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
