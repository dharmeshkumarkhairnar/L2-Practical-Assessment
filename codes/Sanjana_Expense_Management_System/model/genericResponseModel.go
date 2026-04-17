package model

type ErrorMessage struct {
	Key      string `json:"key,omitempty"`
	ErrorMsg string `json:"errMsg,omitempty"`
}

type ErrorAPIResponse struct {
	Message ErrorMessage `json:"message,omitempty"`
	Errors  string       `json:"errors,omitempty"`
}

type JWT struct {
	AccessSecreteKey  string
	RefreshSecreteKey string
}
