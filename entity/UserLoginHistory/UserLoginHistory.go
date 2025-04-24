package UserLoginHistory

import (
	"time"
)

type UserLoginHistory struct {
	Uuid     string
	UserUuid string
	Date     time.Time
	Ip       string
}
