package model

type BFFLoginUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type BFFLoginUserResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type BFFLogoutUserResponse struct {
	Message string `json:"message"`
}

type BFFAvailableSlotsReponse struct {
	SlotNumbers []int `json:"slotNumbers"`
}

type BFFUserBookingsResponse struct {
	Id            int    `json:"bookingId"` 
	SlotNumber    int    `json:"slotNumber"`
	VehicleNumber string `json:"vehicleNumber"`
	Status        string `json:"status"`
}

type BFFCancelBookingsResponse struct {
	Message string `json:"message"`
}
