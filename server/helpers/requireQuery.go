package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequireQuery(c *gin.Context, key string) (string, bool) {
	value := c.Query(key)
	if value == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("%s required", key),
		})
		return "", false
	}
	return value, true
}

func RequireQueries(c *gin.Context, keys ...string) (map[string]string, bool) {
	values := make(map[string]string)
	for _, key := range keys {
		v := c.Query(key)
		if v == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": fmt.Sprintf("%s required", key),
			})
			return nil, false
		}
		values[key] = v
	}
	return values, true
}
