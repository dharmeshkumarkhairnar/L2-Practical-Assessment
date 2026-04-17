package model

import "time"

type BFFAddExpenseRequest struct {
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
}

type BFFAddExpenseResponse struct {
	Id          int64
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	Date        time.Time `json:"date"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
