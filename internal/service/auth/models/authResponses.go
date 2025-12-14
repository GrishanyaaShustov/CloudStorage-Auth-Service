package models

import "time"

// RegisterResponse returns information about new registered(created) user
type RegisterResponse struct {
	UserID string
}

// LoginResponse returns only access-token to set it in HttpOnly secure cookie bcs refresh-token sets in redis in auth-service
type LoginResponse struct {
	AccessToken string
}

// RefreshAccessResponse returns access-token to update old and set it in HttpOnly secure cookie
type RefreshAccessResponse struct {
	AccessToken string
}

// LogoutResponse returns nothing bcs logout request delete tokens
type LogoutResponse struct {
}

// UserInformationResponse returns base information about user
type UserInformationResponse struct {
	Email        string
	UserID       string
	RegisterDate time.Time
}
