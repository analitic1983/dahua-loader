package User

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Status string

const (
	StatusActive   Status = "active"
	StatusInActive Status = "inactive"
)

type User struct {
	Uuid      string
	Name      string
	Email     string
	PassHash  *string
	CreatedAt time.Time
	Status    Status
}

func PasswordHash(password string) string {
	hash := sha256.Sum256([]byte(password))
	passHash := hex.EncodeToString(hash[:])
	return passHash
}

func (user *User) SetPassword(newPassword string) {
	passHash := PasswordHash(newPassword)
	user.PassHash = &passHash
}
