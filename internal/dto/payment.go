package dto

import "time"

type Payment struct {
	PaymentID string    `json:"payment_id" validated:"required"`
	Timestamp time.Time `json:"timestamp" validated:"required"`
}
