package v1

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/augustus281/cqrs-pattern/internal/order/aggregate"
	"github.com/augustus281/cqrs-pattern/pkg/es"
)

type CompleteOrder interface {
	Handle(ctx context.Context, command *CompleteOrderCommand) error
}

type completeOrder struct {
	es es.AggregateStore
}

func NewCompleteOrder(es es.AggregateStore) CompleteOrder {
	return &completeOrder{
		es: es,
	}
}

func (c *completeOrder) Handle(ctx context.Context, command *CompleteOrderCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "completeOrder.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	order, err := aggregate.LoadOrderAggregate(ctx, c.es, command.GetAggregateID())
	if err != nil {
		return err
	}

	if err := order.CompleteOrder(ctx, command.DeliveryTimestamp); err != nil {
		return err
	}

	return c.es.Save(ctx, order)
}
