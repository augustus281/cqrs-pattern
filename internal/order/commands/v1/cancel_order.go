package v1

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/augustus281/cqrs-pattern/internal/order/aggregate"
	"github.com/augustus281/cqrs-pattern/pkg/es"
)

type CancelOrder interface {
	Handle(ctx context.Context, command *CancelOrderCommand) error
}

type cancelOrderCommand struct {
	es es.AggregateStore
}

func NewCancelOrder(es es.AggregateStore) *cancelOrderCommand {
	return &cancelOrderCommand{
		es: es,
	}
}

func (c *cancelOrderCommand) Handle(ctx context.Context, command *CancelOrderCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cancelOrderCommand.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.AggregateID))

	order, err := aggregate.LoadOrderAggregate(ctx, c.es, command.GetAggregateID())
	if err != nil {
		return err
	}

	if err := order.CancelOrder(ctx, command.CancelReason); err != nil {
		return err
	}

	return c.es.Save(ctx, order)
}
