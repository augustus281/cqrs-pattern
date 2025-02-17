package mappers

import (
	"github.com/augustus281/cqrs-pattern/internal/dto"
	"github.com/augustus281/cqrs-pattern/internal/order/models"
)

func PaymentResponseFromModel(payment models.Payment) dto.Payment {
	return dto.Payment{
		PaymentID: payment.PaymentID,
		Timestamp: payment.Timestamp,
	}
}
