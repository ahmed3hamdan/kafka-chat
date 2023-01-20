package api

type GetUserByUsernameRequestBody struct {
	Username string `json:"username" validate:"required,max=20,username"`
}

type GetUserByUsernameResponse struct {
	UserID   int64  `json:"userID"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
