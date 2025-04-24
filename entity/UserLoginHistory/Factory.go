package UserLoginHistory

import (
	"koshmin/dahua-loader/common"
	"koshmin/dahua-loader/entity/User"
	"time"
)

func Create(user User.User, ip string) UserLoginHistory {
	var userLoginHistory UserLoginHistory
	userLoginHistory.Uuid = common.Uuid7String()
	userLoginHistory.UserUuid = user.Uuid
	userLoginHistory.Ip = ip
	userLoginHistory.Date = time.Now()
	return userLoginHistory
}
