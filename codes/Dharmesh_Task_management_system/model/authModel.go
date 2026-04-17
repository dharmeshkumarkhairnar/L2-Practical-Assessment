package model

type BFFLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=20,checkPassword"`
}

type BFFLoginResponse struct {
	Status string `json:"status" example:"success or failed"`
	Token  string `json:"token"`
}

type BFFLogoutResponse struct {
	Status string `json:"status" example:"success or failed"`
}
