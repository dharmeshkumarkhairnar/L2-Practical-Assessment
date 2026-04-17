package model

type BFFPutExpenseRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

type BFFPutExpenseResponse struct {
	Message string `json:"message"`
}
