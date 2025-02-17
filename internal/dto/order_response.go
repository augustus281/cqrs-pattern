package dto

import "time"

type OrderResponseDto struct {
	ID              string     `json:"id,omitempty"`
	OrderID         string     `json:"order_it,omitempty"`
	ShopItems       []ShopItem `json:"shop_items,omitempty"`
	AccountEmail    string     `json:"account_email,omitempty" validate:"required,email"`
	DeliveryAddress string     `json:"delivery_address,omitempty"`
	CancelReason    string     `json:"cancel_reason,omitempty"`
	TotalPrice      float64    `json:"total_price,omitempty"`
	DeliveredTime   time.Time  `json:"delivered_time,omitempty"`
	Created         bool       `json:"created,omitempty"`
	Paid            bool       `json:"paid,omitempty"`
	Submitted       bool       `json:"submitted,omitempty"`
	Completed       bool       `json:"completed,omitempty"`
	Canceled        bool       `json:"canceled,omitempty"`
	Payment         Payment    `json:"payment,omitempty"`
}
