package v1

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/augustus281/cqrs-pattern/internal/order/aggregate"
	"github.com/augustus281/cqrs-pattern/pkg/es"
)

type PayOrder interface {
	Handle(ctx context.Context, command *PayOrderCommand) error
}

type payOrder struct {
	es es.AggregateStore
}

func NewPayOrder(es es.AggregateStore) PayOrder {
	return &payOrder{
		es: es,
	}
}

func (c *payOrder) Handle(ctx context.Context, command *PayOrderCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "payOrder.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	order, err := aggregate.LoadOrderAggregate(ctx, c.es, command.GetAggregateID())
	if err != nil {
		return err
	}

	if err := order.PayOrder(ctx, command.Payment); err != nil {
		return err
	}

	return c.es.Save(ctx, order)
}
