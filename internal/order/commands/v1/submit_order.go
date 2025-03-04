package v1

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/augustus281/cqrs-pattern/internal/order/aggregate"
	"github.com/augustus281/cqrs-pattern/pkg/es"
)

type SubmitOrder interface {
	Handle(ctx context.Context, command *SubmitOrderCommand) error
}

type submitOrder struct {
	es es.AggregateStore
}

func NewSubmitOrder(es es.AggregateStore) SubmitOrder {
	return &submitOrder{
		es: es,
	}
}

func (c *submitOrder) Handle(ctx context.Context, command *SubmitOrderCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "submitOrder.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	order, err := aggregate.LoadOrderAggregate(ctx, c.es, command.GetAggregateID())
	if err != nil {
		return err
	}

	if err := order.SubmitOrder(ctx); err != nil {
		return err
	}

	return c.es.Save(ctx, order)
}
