package helpers

import "github.com/google/uuid"

func CreateToken() string {
	return uuid.New().String()
}
