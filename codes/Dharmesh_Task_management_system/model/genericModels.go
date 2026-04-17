package model

type ErrorMessage struct {
	Key     string `json:"key,omitempty"`
	Message string `json:"message,omitempty"`
}

type ErrorAPIResponse struct {
	ErrorMsg ErrorMessage `json:"errorMsg,omitempty"`
	Message  string       `json:"message,omitempty"`
}
