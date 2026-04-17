package model

type BFFLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,passwordFormat"`
}

type BFFLoginResponse struct {
	Message string `json:"message" example:"Login Successful"`
	Token   string `json:"token"`
}
