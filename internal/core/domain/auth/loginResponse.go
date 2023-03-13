package auth

type LoginResponse struct {
	Email       string `json:"email"`
	TokenString string `json:"token"`
}
