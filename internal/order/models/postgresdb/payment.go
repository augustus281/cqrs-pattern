package postgresdb

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	orderservice "github.com/augustus281/cqrs-pattern/api"
)

type Payment struct {
	PaymentID string    `json:"payment_id,omitempty" validate:"required"`
	Timestamp time.Time `json:"timestamp,omitempty" validate:"required"`
}

func PaymentToProto(payment Payment) *orderservice.Payment {
	return &orderservice.Payment{
		Id:        payment.PaymentID,
		Timestamp: timestamppb.New(payment.Timestamp),
	}
}

func (p *Payment) String() string {
	return fmt.Sprintf("PaymentID: {%s}, Timestamp: {%s}", p.PaymentID, p.Timestamp.UTC().String())
}
