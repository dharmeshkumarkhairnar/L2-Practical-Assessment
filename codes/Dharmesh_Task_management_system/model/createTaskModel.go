package model

type BFFCreateTaskRequest struct {
	Title       string `json:"title" validate:"required,max=200"`
	Description string `json:"description,omitempty" validate:"max=1000"`
	Status      string `json:"status,omitempty" validate:"checkStatus"`
	Priority    string `json:"priority" validate:"required,checkPriority"`
}

type BFFCreateTaskResponse struct {
	Message string `json:"message"`
}
