package model

type BFFDelExpenseRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

type BFFDelExpenseResponse struct {
	Message string `json:"message"`
}
