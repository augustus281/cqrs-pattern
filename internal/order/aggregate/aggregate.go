package aggregate

import (
	"github.com/pkg/errors"

	v1 "github.com/augustus281/cqrs-pattern/internal/order/events/v1"
	"github.com/augustus281/cqrs-pattern/internal/order/models"
	"github.com/augustus281/cqrs-pattern/pkg/es"
)

const (
	OrderAggregateType es.AggregateType = "order"
)

type OrderAggregate struct {
	*es.AggregateBase
	Order *models.Order
}

func NewOrderAggregateWithID(id string) *OrderAggregate {
	if id == "" {
		return nil
	}

	aggregate := NewOrderAggregate()
	aggregate.SetID(id)
	aggregate.Order.ID = id
	return aggregate
}

func NewOrderAggregate() *OrderAggregate {
	orderAggregate := &OrderAggregate{
		Order: models.NewOrder(),
	}
	base := es.NewAggregateBase(orderAggregate.When)
	base.SetType(OrderAggregateType)
	orderAggregate.AggregateBase = base
	return orderAggregate
}

func (a *OrderAggregate) When(event es.Event) error {
	switch event.GetEventType() {
	case v1.OrderCreated:
		return a.onOrderCreated(event)
	case v1.OrderPaid:
		return a.onOrderPaid(event)
	case v1.OrderSubmitted:
		return a.onOrderSubmitted()
	case v1.OrderCanceled:
		return a.onOrderCanceled(event)
	case v1.OrderCompleted:
		return a.onOrderCompleted(event)
	case v1.ShoppingCartUpdated:
		return a.onShoppingCartUpdated(event)
	case v1.DeliveryAddressChanged:
		return a.onChangeDeliveryAddress(event)
	default:
		return es.ErrInvalidEventType
	}
}

func (a *OrderAggregate) onOrderCreated(event es.Event) error {
	var eventData v1.OrderCreatedEvent
	if err := event.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJSONData")
	}
	a.Order.AccountEmail = eventData.AccountEmail
	a.Order.ShopItems = eventData.ShopItems
	a.Order.TotalPrice = GetShopItemsTotalPrice(eventData.ShopItems)
	a.Order.DeliveryAddress = eventData.DeliveryAddress
	return nil
}

func (a *OrderAggregate) onOrderPaid(event es.Event) error {
	var payment models.Payment
	if err := event.GetJsonData(&payment); err != nil {
		return errors.Wrap(err, "GetJSONData")
	}
	a.Order.Paid = true
	a.Order.Payment = payment
	return nil
}

func (a *OrderAggregate) onOrderSubmitted() error {
	a.Order.Submitted = true
	return nil
}

func (a *OrderAggregate) onOrderCompleted(event es.Event) error {
	var eventData v1.OrderCompletedEvent
	if err := event.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJSONData")
	}

	a.Order.Completed = true
	a.Order.DeliveredTime = eventData.DeliveryTimestamp
	a.Order.Canceled = false
	return nil
}

func (a *OrderAggregate) onOrderCanceled(event es.Event) error {
	var eventData v1.OrderCanceledEvent
	if err := event.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJSONData")
	}
	a.Order.Canceled = true
	a.Order.Completed = false
	a.Order.CancelReason = eventData.CancelReason
	return nil
}

func (a *OrderAggregate) onShoppingCartUpdated(event es.Event) error {
	var eventData v1.ShoppingCartUpdatedEvent
	if err := event.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJSONData")
	}
	a.Order.ShopItems = eventData.ShopItems
	a.Order.TotalPrice = GetShopItemsTotalPrice(eventData.ShopItems)
	return nil
}

func (a *OrderAggregate) onChangeDeliveryAddress(event es.Event) error {
	var eventData v1.OrderDeliveryAddressChangedEvent
	if err := event.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJSONData")
	}
	a.Order.DeliveryAddress = eventData.DeliveryAddress
	return nil
}
