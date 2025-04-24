package User

import (
	"koshmin/dahua-loader/common"
	"koshmin/dahua-loader/common/validators"
	"koshmin/dahua-loader/database"
)

func (user *User) Validate() common.ValidatorErrors {
	var existingUser User
	gorm := database.GormDB

	validatorErrors := common.ValidatorErrors{}

	if user.Email == "" {
		// Check for empty
		validatorErrors.Add("email", "Email should be not empty")
	} else if !validators.IsEmailValid(user.Email) {
		// Invalid email format
		validatorErrors.Add("email", "Invalid email format")
	} else {
		// Check for email already exists
		gormResult := gorm.Where("email = ? AND status = ?", user.Email, StatusActive).First(&existingUser)
		if gormResult.Error == nil {
			validatorErrors.Add("email", "Email already exists for an active user")
		}
	}
	return validatorErrors
}
