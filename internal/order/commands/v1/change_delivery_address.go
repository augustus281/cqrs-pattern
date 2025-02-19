package v1

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/augustus281/cqrs-pattern/internal/order/aggregate"
	"github.com/augustus281/cqrs-pattern/pkg/es"
)

type ChangeDeliveryAddress interface {
	Handle(ctx context.Context, command *ChangeDeliveryAddressCommand) error
}

type changeDeliveryAddress struct {
	es es.AggregateStore
}

func NewChangeDeliveryAddress(es es.AggregateStore) *changeDeliveryAddress {
	return &changeDeliveryAddress{
		es: es,
	}
}

func (c *changeDeliveryAddress) Handle(ctx context.Context, command *ChangeDeliveryAddressCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "changeDeliveryAddress.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	order, err := aggregate.LoadOrderAggregate(ctx, c.es, command.GetAggregateID())
	if err != nil {
		return err
	}

	if err := order.ChangeDeliveryAddress(ctx, command.DeliveryAddress); err != nil {
		return err
	}

	return c.es.Save(ctx, order)
}
