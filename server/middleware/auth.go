package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"koshmin/dahua-loader/database"
	"koshmin/dahua-loader/entity/User"
	"net/http"
	"strings"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.FullPath() == "/auth" || c.FullPath() == "/" {
			// Skip logging for this path
			c.Next()
			return
		}
		// Check login
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{
				"success": false,
				"message": "No \"Authorization\" header specified",
			})
			time.Sleep(500 * time.Millisecond)
			c.Abort()
			return
		}
		// Cat "Bearer " prefix
		token = strings.TrimPrefix(token, "Bearer ")

		gorm := database.GormDB
		redis := database.RedisDB
		ctx := context.Background()

		userTokenInfo := redis.Get(ctx, "UserToken["+token+"]")
		userUuid := userTokenInfo.Val()

		var existingUser User.User
		gormResult := gorm.Where("uuid = ? AND status = ?", userUuid, User.StatusActive).First(&existingUser)
		if gormResult.Error != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "Invalid token or user status",
			})
			time.Sleep(500 * time.Millisecond)
			c.Abort()
			return
		}

		c.Next()
	}
}
