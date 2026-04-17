package model

type BFFUpdateTaskRequest struct {
	Title       string `json:"title" validate:"required,max=200"`
	Description string `json:"description,omitempty" validate:"max=1000"`
	Status      string `json:"status,omitempty"`
	Priority    string `json:"priority" validate:"required"`
}
type BFFUpdateTaskResponse struct {
	Message string `json:"message"`
}
