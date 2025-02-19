package v1

import (
	"github.com/augustus281/cqrs-pattern/internal/order/models"
	"github.com/augustus281/cqrs-pattern/pkg/es"
	"time"
)

type CreateOrderCommand struct {
	es.BaseCommand
	ShopItems       []*models.ShopItem `json:"shop_items" validate:"required"`
	AccountEmail    string             `json:"account_email" validate:"required,email"`
	DeliveryAddress string             `json:"delivery_address" validate:"required"`
}

func NewCreateOrderCommand(aggregateID string, shopItems []*models.ShopItem, accountEmail, deliveryAddress string) *CreateOrderCommand {
	return &CreateOrderCommand{
		BaseCommand:     es.NewBaseCommand(aggregateID),
		ShopItems:       shopItems,
		AccountEmail:    accountEmail,
		DeliveryAddress: deliveryAddress,
	}
}

type PayOrderCommand struct {
	models.Payment
	es.BaseCommand
}

func NewPayOrderCommand(payment models.Payment, aggregateID string) *PayOrderCommand {
	return &PayOrderCommand{
		Payment:     payment,
		BaseCommand: es.NewBaseCommand(aggregateID),
	}
}

type SubmitOrderCommand struct {
	es.BaseCommand
}

func NewSubmitOrderCommand(aggregateID string) *SubmitOrderCommand {
	return &SubmitOrderCommand{
		es.NewBaseCommand(aggregateID),
	}
}

type UpdateShoppingCartCommand struct {
	es.BaseCommand
	ShopItems []*models.ShopItem `json:"shop_items" validate:"required"`
}

func NewUpdateShoppingCartCommand(aggregateID string, shopItems []*models.ShopItem) *UpdateShoppingCartCommand {
	return &UpdateShoppingCartCommand{
		BaseCommand: es.NewBaseCommand(aggregateID),
		ShopItems:   shopItems,
	}
}

type CancelOrderCommand struct {
	es.BaseCommand
	CancelReason string `json:"cancel_reason" validate:"required"`
}

func NewCancelOrderCommand(aggregateID, cancelReason string) *CancelOrderCommand {
	return &CancelOrderCommand{
		BaseCommand:  es.NewBaseCommand(aggregateID),
		CancelReason: cancelReason,
	}
}

type CompleteOrderCommand struct {
	es.BaseCommand
	DeliveryTimestamp time.Time `json:"delivery_timestamp" validate:"required"`
}

func NewCompleteOrderCommand(aggregateID string, deliveryTimestamp time.Time) *CompleteOrderCommand {
	return &CompleteOrderCommand{
		BaseCommand:       es.NewBaseCommand(aggregateID),
		DeliveryTimestamp: deliveryTimestamp,
	}
}

type ChangeDeliveryAddressCommand struct {
	es.BaseCommand
	DeliveryAddress string `json:"delivery_address" validate:"required"`
}

func NewChangeDeliveryAddressCommand(aggregateID, deliveryAddress string) *ChangeDeliveryAddressCommand {
	return &ChangeDeliveryAddressCommand{
		BaseCommand:     es.NewBaseCommand(aggregateID),
		DeliveryAddress: deliveryAddress,
	}
}
