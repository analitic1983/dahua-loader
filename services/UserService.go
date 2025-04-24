package services

import (
	"fmt"
	"koshmin/dahua-loader/database"
	"koshmin/dahua-loader/entity/User"
)

func AddAdmin(email string, password string) (User.User, error) {

	user := User.Create("Administrator", email)
	user.SetPassword(password)
	validatorResult := user.Validate()
	if validatorResult.HasErrors() {
		return User.User{}, validatorResult
	}
	gorm := database.GormDB
	result := gorm.Create(&user)
	if result.Error != nil {
		return User.User{}, fmt.Errorf("Db failed: %w", result.Error)
	}
	return user, nil
}
