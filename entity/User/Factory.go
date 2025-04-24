package User

import (
	"koshmin/dahua-loader/common"
	"time"
)

func Create(name string, email string) User {
	user := User{
		Uuid:      common.Uuid7String(),
		Name:      name,
		Email:     email,
		PassHash:  nil,
		CreatedAt: time.Now(),
		Status:    StatusActive,
	}
	return user
}
