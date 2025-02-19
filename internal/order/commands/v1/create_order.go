package v1

import (
	"context"

	"github.com/augustus281/cqrs-pattern/internal/order/aggregate"
	"github.com/augustus281/cqrs-pattern/pkg/es"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type CreateOrder interface {
	Handle(ctx context.Context, command *CreateOrderCommand) error
}

type createOrder struct {
	es es.AggregateStore
}

func NewCreateOrder(es es.AggregateStore) CreateOrder {
	return &createOrder{
		es: es,
	}
}

func (c *createOrder) Handle(ctx context.Context, command *CreateOrderCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "createOrder.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	order, err := aggregate.LoadOrderAggregate(ctx, c.es, command.GetAggregateID())
	if err != nil {
		return err
	}

	if err := order.CreateOrder(ctx, command.ShopItems, command.AccountEmail, command.DeliveryAddress); err != nil {
		return err
	}

	return c.es.Save(ctx, order)
}
