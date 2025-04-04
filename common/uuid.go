package common

import "github.com/google/uuid"

func Uuid4String() string {
	uuid4, _ := uuid.NewRandom()
	return uuid4.String()
}

func Uuid7String() string {
	uuid7, _ := uuid.NewV7()
	return uuid7.String()
}
