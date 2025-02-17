package v1

import (
	"time"

	"github.com/augustus281/cqrs-pattern/internal/order/models"
	"github.com/augustus281/cqrs-pattern/pkg/es"
)

const (
	OrderCreated           = "V1_ORDER_CREATED"
	OrderPaid              = "V1_ORDER_PAID"
	OrderSubmitted         = "V1_ORDER_SUBMITTED"
	OrderCompleted         = "V1_ORDER_COMPLETED"
	OrderCanceled          = "V1_ORDER_CANCELED"
	ShoppingCartUpdated    = "V1_SHOPPING_CART_UPDATED"
	DeliveryAddressChanged = "V1_DELIVERY_ADDRESS_CHANGED"
)

type OrderCreatedEvent struct {
	ShopItems       []*models.ShopItem `json:"shop_items"`
	AccountEmail    string             `json:"account_email"`
	DeliveryAddress string             `json:"delivery_address"`
}

func NewOrderCreatedEvent(aggregate es.Aggregate, shopItems []*models.ShopItem, accountEmail, deliveryAddress string) (es.Event, error) {
	eventData := OrderCreatedEvent{
		ShopItems:       shopItems,
		AccountEmail:    accountEmail,
		DeliveryAddress: deliveryAddress,
	}
	event := es.NewBaseEvent(aggregate, OrderCreated)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}

func NewOrderPaidEvent(aggregate es.Aggregate, payment *models.Payment) (es.Event, error) {
	event := es.NewBaseEvent(aggregate, OrderPaid)
	if err := event.SetJsonData(&payment); err != nil {
		return es.Event{}, err
	}
	return event, nil
}

func NewSubmitOrderEvent(aggregate es.Aggregate) (es.Event, error) {
	return es.NewBaseEvent(aggregate, OrderSubmitted), nil
}

type ShoppingCartUpdatedEvent struct {
	ShopItems []*models.ShopItem `json:"shop_items"`
}

func NewShoppingCartUpdateEvent(aggregate es.Aggregate, shopItems []*models.ShopItem) (es.Event, error) {
	eventData := ShoppingCartUpdatedEvent{
		ShopItems: shopItems,
	}
	event := es.NewBaseEvent(aggregate, ShoppingCartUpdated)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, nil
	}
	return event, nil
}

type OrderDeliveryAddressChangedEvent struct {
	DeliveryAddress string `json:"delivery_address,omitempty"`
}

func NewDeliveryAddressChangedEvent(aggregate es.Aggregate, deliveryAddress string) (es.Event, error) {
	eventData := OrderDeliveryAddressChangedEvent{
		DeliveryAddress: deliveryAddress,
	}
	event := es.NewBaseEvent(aggregate, DeliveryAddressChanged)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}

type OrderCanceledEvent struct {
	CancelReason string `json:"cancel_reason"`
}

func NewOrderCanceledEvent(aggregate es.Aggregate, cancelReason string) (es.Event, error) {
	eventData := OrderCanceledEvent{
		CancelReason: cancelReason,
	}
	event := es.NewBaseEvent(aggregate, OrderCanceled)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, nil
	}
	return event, nil
}

type OrderCompletedEvent struct {
	DeliveryTimestamp time.Time `json:"delivery_timestamp"`
}

func NewOrderCompletedEvent(aggregate es.Aggregate, deliveryTime time.Time) (es.Event, error) {
	eventData := OrderCompletedEvent{
		DeliveryTimestamp: deliveryTime,
	}
	event := es.NewBaseEvent(aggregate, OrderCompleted)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, nil
	}
	return event, nil
}
