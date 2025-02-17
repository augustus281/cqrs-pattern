package repository

import (
	"context"
	"encoding/json"
	"fmt"

	v7 "github.com/olivere/elastic/v7"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"

	"github.com/augustus281/cqrs-pattern/global"
	"github.com/augustus281/cqrs-pattern/internal/dto"
	"github.com/augustus281/cqrs-pattern/internal/mappers"
	"github.com/augustus281/cqrs-pattern/internal/order/models"
	"github.com/augustus281/cqrs-pattern/pkg/tracing"
	"github.com/augustus281/cqrs-pattern/pkg/utils"
)

const (
	shopItemTitle            = "shopItems.title"
	shopItemDescription      = "shopItems.description"
	minimumNumberShouldMatch = 1
)

type elasticRepository struct {
	elasticClient *v7.Client
}

func NewElasticRepository(elasticClient *v7.Client) *elasticRepository {
	return &elasticRepository{
		elasticClient: elasticClient,
	}
}

func (e *elasticRepository) IndexOrder(ctx context.Context, order *models.OrderProjection) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "elasticRepository.IndexOrder")
	defer span.Finish()
	span.LogFields(log.String("OrderID", order.OrderID))

	res, err := e.elasticClient.Index().
		Index(global.Config.ElasticIndexes.Orders).
		BodyJson(order).
		Id(order.OrderID).
		Do(ctx)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "elasticClient.Index")
	}

	global.Logger.Debug(fmt.Sprintf("(IndexOrder) result: {%s}", res.Result))
	return nil
}

func (e *elasticRepository) GetByID(ctx context.Context, orderID string) (*models.OrderProjection, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "elasticRepository.GetByID")
	defer span.Finish()
	span.LogFields(log.String("OrderID", orderID))

	result, err := e.elasticClient.Get().
		Index(global.Config.ElasticIndexes.Orders).
		Index(orderID).
		FetchSource(true).
		Do(ctx)
	if err != nil {
		tracing.TraceErr(span, err)
		return nil, errors.Wrap(err, "elasticClient.Get")
	}

	jsonData, err := result.Source.MarshalJSON()
	if err != nil {
		tracing.TraceErr(span, err)
		return nil, errors.Wrap(err, "Source.MarshalJSON")
	}

	var order models.OrderProjection
	if err := json.Unmarshal(jsonData, &order); err != nil {
		tracing.TraceErr(span, err)
		return nil, errors.Wrap(err, "json.Unmarshal")
	}

	return &order, nil
}

func (e *elasticRepository) UpdateOrder(ctx context.Context, order *models.OrderProjection) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "elasticRepository.UpdateShoppingCart")
	defer span.Finish()
	span.LogFields(log.String("OrderID", order.OrderID))

	res, err := e.elasticClient.Update().
		Index(global.Config.ElasticIndexes.Orders).
		Id(order.OrderID).
		Doc(order).
		FetchSource(true).
		Do(ctx)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "elasticClient.Update")
	}
	global.Logger.Debug(fmt.Sprintf("(UpdateShoppingCart) result: {%s}", res.Result))
	return nil
}

func (e *elasticRepository) Search(ctx context.Context, text string, pq *utils.Pagination) (*dto.OrderSearchResponseDTO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "elasticRepository.Search")
	defer span.Finish()
	span.LogFields(log.String("Search", text))

	shouldMatch := v7.NewBoolQuery().
		Should(v7.NewMatchPhrasePrefixQuery(shopItemTitle, text), v7.NewMatchPhrasePrefixQuery(shopItemDescription, text)).
		MinimumNumberShouldMatch(minimumNumberShouldMatch)

	searchResult, err := e.elasticClient.Search(global.Config.ElasticIndexes.Orders).
		Query(shouldMatch).
		From(pq.GetOffset()).
		Explain(global.Config.ElasticSearch.Explain).
		FetchSource(global.Config.ElasticSearch.FetchSource).
		Version(global.Config.ElasticSearch.Version).
		Size(pq.GetSize()).
		Pretty(global.Config.ElasticSearch.Pretty).
		Do(ctx)
	if err != nil {
		tracing.TraceErr(span, err)
		return nil, errors.Wrapf(err, "elasticClient.Search")
	}

	orders := make([]*models.OrderProjection, 0, len(searchResult.Hits.Hits))
	for _, hit := range searchResult.Hits.Hits {
		jsonBytes, err := hit.Source.MarshalJSON()
		if err != nil {
			tracing.TraceErr(span, err)
			return nil, errors.Wrap(err, "Source.MarshalJSON")
		}
		var order models.OrderProjection
		if err := json.Unmarshal(jsonBytes, &order); err != nil {
			tracing.TraceErr(span, err)
			return nil, errors.Wrap(err, "json.Unmarshal")
		}
		orders = append(orders, &order)
	}

	return &dto.OrderSearchResponseDTO{
		Pagination: dto.Pagination{
			TotalCount: searchResult.TotalHits(),
			TotalPages: int64(pq.GetTotalPages(int(searchResult.TotalHits()))),
			Page:       int64(pq.GetPage()),
			Size:       int64(pq.GetSize()),
			HasMore:    pq.GetHasMore(int(searchResult.TotalHits())),
		},
		Orders: mappers.OrdersFromProjections(orders),
	}, nil
}
