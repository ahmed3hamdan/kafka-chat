package api

type RegisterRequestBody struct {
	Name     string `json:"name" validate:"required,max=60"`
	Username string `json:"username" validate:"required,max=20,username"`
	Password string `json:"password" validate:"required,max=72"`
}

type AuthResponse struct {
	UserID int64  `json:"userID"`
	Token  string `json:"token"`
}

type LoginRequestBody struct {
	Username string `json:"username" validate:"required,max=20,username"`
	Password string `json:"password" validate:"required,max=72"`
}
