package auth

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type meResponse struct {
	UserID       string `json:"user_id"`
	Email        string `json:"email"`
	RegisterDate string `json:"register_date"`
}
