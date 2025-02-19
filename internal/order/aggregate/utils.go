package aggregate

import (
	"context"
	"errors"
	"strings"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/augustus281/cqrs-pattern/internal/order/models"
	"github.com/augustus281/cqrs-pattern/pkg/es"
)

func GetShopItemsTotalPrice(shopItems []*models.ShopItem) float64 {
	var totalPrice float64 = 0
	for _, shopItem := range shopItems {
		totalPrice += shopItem.Price * float64(shopItem.Quantity)
	}
	return totalPrice
}

func GetOrderAggregateID(eventAggregateID string) string {
	return strings.ReplaceAll(eventAggregateID, "order-", "")
}

func IsAggregateNotFound(aggregate es.Aggregate) bool {
	return aggregate.GetVersion() == 0
}

func LoadOrderAggregate(ctx context.Context, eventStore es.AggregateStore, aggregateID string) (*OrderAggregate, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LoadOrderAggregate")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", aggregateID))

	order := NewOrderAggregateWithID(aggregateID)

	err := eventStore.Exists(ctx, order.GetID())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return nil, err
	}

	if err := eventStore.Load(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}
