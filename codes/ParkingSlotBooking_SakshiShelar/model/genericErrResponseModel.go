package model

type ErrorMessage struct {
	Key      string `json:"key"`
	ErrorMsg string `json:"errorMessage"`
}

type ErrorMessageResponse struct {
	Message      ErrorMessage `json:"message"`
	ErrorMessage string       `json:"error"`
}

type GenericErrorMsgResponse struct {
	Message string `json:"message"`
}
