package models

import (
	"fmt"
	"time"

	orderservice "github.com/augustus281/cqrs-pattern/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Order struct {
	ID              string      `json:"id,omitempty"`
	ShopItems       []*ShopItem `json:"shop_items,omitempty"`
	AccountEmail    string      `json:"account_email,omitempty"`
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

func (o *Order) String() string {
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

func NewOrder() *Order {
	return &Order{
		ShopItems: make([]*ShopItem, 0),
		Paid:      false,
		Submitted: false,
		Completed: false,
		Canceled:  false,
	}
}

func OrderToProto(order *Order, id string) *orderservice.Order {
	return &orderservice.Order{
		Id:                id,
		ShopItems:         ShopItemsToProto(order.ShopItems),
		Paid:              order.Paid,
		Submitted:         order.Submitted,
		Completed:         order.Completed,
		Canceled:          order.Canceled,
		CancelReason:      order.CancelReason,
		DeliveryTimestamp: timestamppb.New(order.DeliveredTime),
		DeliveryAddress:   order.DeliveryAddress,
		AccountEmail:      order.AccountEmail,
		TotalPrice:        order.TotalPrice,
		Payment:           PaymentToProto(order.Payment),
	}
}
