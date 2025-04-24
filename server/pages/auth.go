package pages

import (
	"context"
	"github.com/gin-gonic/gin"
	"koshmin/dahua-loader/common"
	"koshmin/dahua-loader/database"
	"koshmin/dahua-loader/entity/User"
	"koshmin/dahua-loader/entity/UserLoginHistory"
	"koshmin/dahua-loader/server/helpers"
	"net/http"
	"time"
)

func Auth(c *gin.Context) {
	var jsonData map[string]interface{}

	if err := c.BindJSON(&jsonData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	email, _ := jsonData["email"].(string)
	password, _ := jsonData["password"].(string)

	// Simplest brute force protection
	time.Sleep(300 * time.Millisecond)

	ctx := context.Background()
	gorm := database.GormDB
	redis := database.RedisDB

	var existingUser User.User
	gormResult := gorm.Where("email = ? AND pass_hash = ? AND status = ?", email, User.PasswordHash(password), User.StatusActive).First(&existingUser)
	if gormResult.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Invalid email or password",
		})
		return
	}

	saveLoginHistory(existingUser, c.ClientIP())

	token := common.Uuid4String()
	err := redis.Set(ctx, "UserToken["+token+"]", existingUser.Uuid, 10*time.Hour).Err()
	if err != nil {
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"message": "Successfully logged in",
		"token":   token,
	})
}

func saveLoginHistory(user User.User, ip string) {
	gorm := database.GormDB

	userLoginHistory := UserLoginHistory.Create(user, ip)
	result := gorm.Create(&userLoginHistory)

	if result.Error != nil {
		panic(result.Error)
	}
}
