package queries

import "github.com/augustus281/cqrs-pattern/pkg/utils"

type OrderQueries struct {
	GetOrderByID GetOrderByIdQueryHandler
	SearchOrders SearchOrdersQueryHandler
}

func NewOrderQueries(getOrderByID GetOrderByIdQueryHandler, searchOrders SearchOrdersQueryHandler) *OrderQueries {
	return &OrderQueries{
		GetOrderByID: getOrderByID,
		SearchOrders: searchOrders,
	}
}

type GetOrderByIDQuery struct {
	ID string
}

func NewGetOrderByIDQuery(id string) *GetOrderByIDQuery {
	return &GetOrderByIDQuery{
		ID: id,
	}
}

type SearchOrdersQuery struct {
	SearchText string `form:"search_text"`
	Pq         *utils.Pagination
}

func NewSearchOrdersQuery(searchText string, pq *utils.Pagination) *SearchOrdersQuery {
	return &SearchOrdersQuery{
		SearchText: searchText,
		Pq:         pq,
	}
}
