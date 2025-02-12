package dto

type CancelOrderRequest struct {
	CancelReason string `json:"cancel_reason" validate:"required"`
}
