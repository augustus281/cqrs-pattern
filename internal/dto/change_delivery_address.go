package dto

type ChangeDeliveryAddressRequest struct {
	DeliveryAddress string `json:"delivery_address" validate:"required"`

}