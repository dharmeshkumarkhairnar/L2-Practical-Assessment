package model

type ErrorAPIResponse struct {
	Message ErrorMessage `json:"message,omitempty"`
	Error   string       `json:"error,omitempty"`
}

type ErrorMessage struct {
	Key          string `json:"key,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type LoginSuccessModel struct {
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}

type LogoutSuccessModel struct {
	Message string `json:"message"`
}

type GetOrderSuccessful struct {
	Message string             `json:"message"`
	Orders  []GetOrderResponse `json:"orders"`
}

type UpdateOrderSuccessful struct {
	Message string `json:"message"`
}

type GetOrderResponse struct {
	ProductName string  `json:"product_name" `
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"`
}

type DeleteOrderResponse struct {
	Message string `json:"message"`
}

type CreateOrderResponse struct {
	Message string `json:"message"`
}
