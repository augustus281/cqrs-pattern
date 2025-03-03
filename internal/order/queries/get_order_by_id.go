package queries

import (
	"context"
	"github.com/augustus281/cqrs-pattern/internal/mappers"
	"github.com/augustus281/cqrs-pattern/internal/order/aggregate"
	"github.com/pkg/errors"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/augustus281/cqrs-pattern/internal/order/models"
	"github.com/augustus281/cqrs-pattern/internal/order/repository"
	"github.com/augustus281/cqrs-pattern/pkg/es"
)

type GetOrderByIdQueryHandler interface {
	Handle(ctx context.Context, query *GetOrderByIDQuery) (*models.OrderProjection, error)
}

type getOrderByIDHandler struct {
	es        es.AggregateStore
	mongoRepo repository.MongoOrderRepository
}

func NewGetOrderByIDHandler(es es.AggregateStore, mongoRepo repository.MongoOrderRepository) *getOrderByIDHandler {
	return &getOrderByIDHandler{
		es:        es,
		mongoRepo: mongoRepo,
	}
}

func (q *getOrderByIDHandler) Handle(ctx context.Context, query *GetOrderByIDQuery) (*models.OrderProjection, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getOrderByIDHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", query.ID))

	orderProjection, err := q.mongoRepo.GetByID(ctx, query.ID)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}
	if orderProjection != nil {
		return orderProjection, nil
	}

	order := aggregate.NewOrderAggregateWithID(query.ID)
	if err := q.es.Load(ctx, order); err != nil {
		return nil, err
	}

	if aggregate.IsAggregateNotFound(order) {
		return nil, aggregate.ErrOrderNotFound
	}

	orderProjection = mappers.OrderProjectionFromAggregate(order)

	_, err = q.mongoRepo.Insert(ctx, orderProjection)
	if err != nil {
		return nil, err
	}

	return orderProjection, nil
}
