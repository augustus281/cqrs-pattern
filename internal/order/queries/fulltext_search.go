package queries

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/augustus281/cqrs-pattern/internal/dto"
	"github.com/augustus281/cqrs-pattern/internal/order/repository"
	"github.com/augustus281/cqrs-pattern/pkg/es"
)

type SearchOrdersQueryHandler interface {
	Handle(ctx context.Context, command *SearchOrdersQuery) (*dto.OrderSearchResponseDTO, error)
}

type searchOrdersHandler struct {
	es          es.AggregateStore
	elasticRepo repository.ElasticOrderRepository
}

func NewSearchOrdersHandler(es es.AggregateStore, elasticRepo repository.ElasticOrderRepository) *searchOrdersHandler {
	return &searchOrdersHandler{
		es:          es,
		elasticRepo: elasticRepo,
	}
}

func (s *searchOrdersHandler) Handle(ctx context.Context, command *SearchOrdersQuery) (*dto.OrderSearchResponseDTO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "searchOrdersHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("searchText", command.SearchText))

	return s.elasticRepo.Search(ctx, command.SearchText, command.Pq)
}
