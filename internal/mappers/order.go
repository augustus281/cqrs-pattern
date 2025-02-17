package mappers

import (
	"github.com/augustus281/cqrs-pattern/internal/dto"
	"github.com/augustus281/cqrs-pattern/internal/order/models"
)

func OrderResponseFromProjection(projection *models.OrderProjection) dto.OrderResponseDto {
	return dto.OrderResponseDto{
		ID:              projection.ID,
		OrderID:         projection.OrderID,
		ShopItems:       ShopItemsResponseFromModels(projection.ShopItems),
		AccountEmail:    projection.AccountEmail,
		DeliveryAddress: projection.DeliveryAddress,
		CancelReason:    projection.CancelReason,
		TotalPrice:      projection.TotalPrice,
		DeliveredTime:   projection.DeliveredTime,
		Paid:            projection.Paid,
		Submitted:       projection.Submitted,
		Completed:       projection.Completed,
		Canceled:        projection.Canceled,
		Payment:         PaymentResponseFromModel(projection.Payment),
	}
}

func OrdersFromProjections(projections []*models.OrderProjection) []dto.OrderResponseDto {
	orders := make([]dto.OrderResponseDto, 0, len(projections))
	for _, projection := range projections {
		orders = append(orders, OrderResponseFromProjection(projection))
	}
	return orders
}
