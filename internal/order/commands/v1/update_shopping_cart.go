package v1

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/augustus281/cqrs-pattern/internal/order/aggregate"
	"github.com/augustus281/cqrs-pattern/pkg/es"
)

type UpdateShoppingCart interface {
	Handle(ctx context.Context, command *UpdateShoppingCartCommand) error
}

type updateShoppingCart struct {
	es es.AggregateStore
}

func NewUpdateShoppingCart(es es.AggregateStore) UpdateShoppingCart {
	return &updateShoppingCart{
		es: es,
	}
}

func (c *updateShoppingCart) Handle(ctx context.Context, command *UpdateShoppingCartCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "updateShoppingCart.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	order, err := aggregate.LoadOrderAggregate(ctx, c.es, command.GetAggregateID())
	if err != nil {
		return err
	}

	if err := order.UpdateShoppingCart(ctx, command.ShopItems); err != nil {
		return err
	}

	return c.es.Save(ctx, order)
}
