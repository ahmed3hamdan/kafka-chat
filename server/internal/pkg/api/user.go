package api

type GetUserByUsernameParams struct {
	Username string `params:"username" validate:"required,max=20,username"`
}

type GetUserByUsernameResponse struct {
	UserID   int64  `json:"userID"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
