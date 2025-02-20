package mongomodels

import (
	"fmt"
	"time"
)

type OrderProjection struct {
	ID              string      `json:"id" bson:"_id,omitempty"`
	OrderID         string      `json:"order_id,omitempty" bson:"order_id,omitempty"`
	ShopItems       []*ShopItem `json:"shop_items,omitempty" bson:"shop_items,omitempty"`
	AccountEmail    string      `json:"account_email,omitempty" bson:"account_email,omitempty" validate:"required,email"`
	DeliveryAddress string      `json:"delivery_address,omitempty" bson:"delivery_address,omitempty"`
	CancelReason    string      `json:"cancel_reason,omitempty" bson:"cancel_reason,omitempty"`
	TotalPrice      float64     `json:"total_price,omitempty" bson:"total_price,omitempty"`
	DeliveredTime   time.Time   `json:"delivered_time,omitempty" bson:"delivered_time,omitempty"`
	Paid            bool        `json:"paid,omitempty" bson:"paid,omitempty"`
	Submitted       bool        `json:"submitted,omitempty" bson:"submitted,omitempty"`
	Completed       bool        `json:"completed,omitempty" bson:"completed,omitempty"`
	Canceled        bool        `json:"canceled,omitempty" bson:"canceled,omitempty"`
	Payment         Payment     `json:"payment,omitempty" bson:"payment,omitempty"`
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
