package domain

import "time"

type User struct {
	UserID       string
	Email        string
	PasswordHash string
	RegisterDate time.Time
}
