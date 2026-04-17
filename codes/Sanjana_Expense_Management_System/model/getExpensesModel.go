package model

import "time"

type BFFGetExpenseRequest struct {
	Name     string `json:"name" validate:"required"`
	Category string `json:"category"`
}

type BFFGetExpenseResponse struct {
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}
