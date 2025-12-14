package models

// RegisterRequest are NOT requires authentication
// It sends to register(create) new user. It not authenticate user
type RegisterRequest struct {
	Email    string
	Password string
}

// LoginRequest are NOT requires authentication
// It sends to authenticate user and set refresh-token in redis and set access-token in HttpOnly secure cookie
type LoginRequest struct {
	Email    string
	Password string
}

// RefreshAccessRequest are requires authentication
// It sends to refresh access-token stored in HtppOnly Secure cookie
type RefreshAccessRequest struct {
	AccessToken string
}

// LogoutRequest are requires authentication
// It sends to get delete refresh-token from redis and access-token from HtppOnly Secure cookie
type LogoutRequest struct {
	AccessToken string
}

// UserInformationRequest are requires authentication
// It sends to get main user information
type UserInformationRequest struct {
	AccessToken string
}
