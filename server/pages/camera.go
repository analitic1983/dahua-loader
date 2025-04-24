package pages

import (
	"github.com/gin-gonic/gin"
	"koshmin/dahua-loader/common/json"
	"koshmin/dahua-loader/database"
	"koshmin/dahua-loader/entity/Camera"
	"net/http"
	"strconv"
)

func CameraList(c *gin.Context) {
	const DefaultLimit = "20"

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", DefaultLimit))
	offset, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))

	gorm := database.GormDB

	var cameras []Camera.Camera
	gormResult := gorm.Where("status = ?", Camera.StatusInActive).Offset(offset).Limit(limit).Find(&cameras)
	if gormResult.Error != nil {
		panic(gormResult.Error)
	}

	cameraListArray := json.MarshalMap(cameras, Camera.CameraSerializeMap)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"list":    cameraListArray,
	})
}
