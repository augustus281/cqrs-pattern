package service

import (
	v1 "github.com/augustus281/cqrs-pattern/internal/order/commands/v1"
	"github.com/augustus281/cqrs-pattern/internal/order/queries"
	"github.com/augustus281/cqrs-pattern/internal/order/repository"
	"github.com/augustus281/cqrs-pattern/pkg/es"
)

type OrderService struct {
	Commands *v1.OrderCommands
	Queries  *queries.OrderQueries
}

func NewOrderService(
	es es.AggregateStore,
	elasticRepository repository.ElasticOrderRepository,
) *OrderService {
	return &OrderService{}
}
