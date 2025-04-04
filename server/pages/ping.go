package pages

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	key := c.Query("key")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"key":     key,
	})
	//time.Sleep(300 * time.Millisecond)
}
