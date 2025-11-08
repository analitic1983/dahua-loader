package pages

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"koshmin/dahua-loader/common/json"
	"koshmin/dahua-loader/database"
	"koshmin/dahua-loader/entity/User"
	"koshmin/dahua-loader/server/helpers"
	"net/http"
	"strconv"
)

func UserList(c *gin.Context) {
	const DefaultLimit = "20"

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", DefaultLimit))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	gorm := database.GormDB

	var users []User.User
	gormResult := gorm.Where("status = ?", User.StatusActive).Offset(offset).Limit(limit).Find(&users)
	if gormResult.Error != nil {
		panic(gormResult.Error)
	}

	// Total
	var usersTotal int64
	gormResult = gorm.Model(&User.User{}).Where("status = ?", User.StatusActive).Count(&usersTotal)
	if gormResult.Error != nil {
		panic(gormResult.Error)
	}

	userList := json.MarshalMap(users, User.UserSerializeMap)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"total":   usersTotal,
		"list":    userList,
	})
}

func UserGet(c *gin.Context) {
	uuid, ok := helpers.RequireQuery(c, "uuid")
	if !ok {
		return
	}

	// Не затеняем пакет gorm — используем локальную переменную db
	db := database.GormDB

	var user User.User
	result := db.
		Where("status = ?", User.StatusActive).
		Where("uuid = ?", uuid).
		First(&user) // First вернёт ErrRecordNotFound, если записи нет

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "User not found",
			})
			return
		}
		panic(result.Error)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user":    json.MarshalMap(user, User.UserSerializeMap),
	})
}
