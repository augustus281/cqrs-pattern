package models

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	orderservice "github.com/augustus281/cqrs-pattern/api"
)

type OrderProjection struct {
	ID              string      `json:"id,omitempty"`
	OrderID         string      `json:"order_id,omitempty"`
	ShopItems       []*ShopItem `json:"shop_items,omitempty"`
	AccountEmail    string      `json:"account_email,,omitempty" validate:"required,email"`
	DeliveryAddress string      `json:"delivery_address,omitempty"`
	CancelReason    string      `json:"cancel_reason,omitempty"`
	TotalPrice      float64     `json:"total_price,omitempty"`
	DeliveredTime   time.Time   `json:"delivered_time,omitempty"`
	Paid            bool        `json:"paid,omitempty"`
	Submitted       bool        `json:"submitted,omitempty"`
	Completed       bool        `json:"completed,omitempty"`
	Canceled        bool        `json:"canceled,omitempty"`
	Payment         Payment     `json:"payment,omitempty"`
}

func (o *OrderProjection) String() string {
	return fmt.Sprintf("ID: {%s}, ShopItems: {%+v}, Paid: {%v}, Submitted: {%v}, "+
		"Completed: {%v}, Canceled: {%v}, CancelReason: {%s}, TotalPrice: {%v}, AccountEmail: {%s}, DeliveryAddress: {%s}, DeliveredTime: {%s}, Payment: {%s}",
		o.ID,
		o.ShopItems,
		o.Paid,
		o.Submitted,
		o.Completed,
		o.Canceled,
		o.CancelReason,
		o.TotalPrice,
		o.AccountEmail,
		o.DeliveryAddress,
		o.DeliveredTime.UTC().String(),
		o.Payment.String(),
	)
}

func OrderProjectionToProto(order *OrderProjection) *orderservice.Order {
	return &orderservice.Order{
		Id:                order.OrderID,
		ShopItems:         ShopItemsToProto(order.ShopItems),
		Paid:              order.Paid,
		Submitted:         order.Submitted,
		Completed:         order.Completed,
		Canceled:          order.Canceled,
		TotalPrice:        order.TotalPrice,
		AccountEmail:      order.AccountEmail,
		CancelReason:      order.CancelReason,
		DeliveryTimestamp: timestamppb.New(order.DeliveredTime),
		DeliveryAddress:   order.DeliveryAddress,
		Payment:           PaymentToProto(order.Payment),
	}
}

func OrderProjectionsToProto(orderProjections []*OrderProjection) []*orderservice.Order {
	orders := make([]*orderservice.Order, 0, len(orderProjections))
	for _, projection := range orderProjections {
		orders = append(orders, OrderProjectionToProto(projection))
	}
	return orders
}

func PaymentFromProto(payment *orderservice.Payment) Payment {
	return Payment{
		PaymentID: payment.GetId(),
		Timestamp: payment.GetTimestamp().AsTime(),
	}
}
