package mongomodels

import (
	"fmt"
	"time"
)

type Payment struct {
	PaymentID string    `json:"payment_id" bson:"payment_id,omitempty" validate:"required"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp,omitempty" validate:"required"`
}

// func PaymentToProto(payment Payment) *orderservice.Payment {
// 	return &orderservice.Payment{
// 		Id:        payment.PaymentID,
// 		Timestamp: timestamppb.New(payment.Timestamp),
// 	}
// }

func (p *Payment) String() string {
	return fmt.Sprintf("PaymentID: {%s}, Timestamp: {%s}", p.PaymentID, p.Timestamp.UTC().String())
}
