package repository

import (
	"context"

	"github.com/augustus281/cqrs-pattern/internal/dto"
	"github.com/augustus281/cqrs-pattern/internal/order/models"
	mongomodels "github.com/augustus281/cqrs-pattern/internal/order/models/mongo_models"
	"github.com/augustus281/cqrs-pattern/pkg/utils"
)

type MongoOrderRepository interface {
	Insert(ctx context.Context, order *models.OrderProjection) (string, error)
	GetByID(ctx context.Context, orderID string) (*models.OrderProjection, error)
	UpdateOrder(ctx context.Context, order *mongomodels.OrderProjection) error

	UpdateCancel(ctx context.Context, order *mongomodels.OrderProjection)
	UpdatePayment(ctx context.Context, order *mongomodels.OrderProjection) error
	Complete(ctx context.Context, order *mongomodels.OrderProjection)
	UpdateDeliveryAddress(ctx context.Context, order *mongomodels.OrderProjection) error
	UpdateSubmit(ctx context.Context, order *mongomodels.OrderProjection) error
}

type ElasticOrderRepository interface {
	IndexOrder(ctx context.Context, order *models.OrderProjection) error
	GetByID(ctx context.Context, orderID string) (*models.OrderProjection, error)
	UpdateOrder(ctx context.Context, order *models.OrderProjection) error
	Search(ctx context.Context, text string, pq *utils.Pagination) (*dto.OrderSearchResponseDTO, error)
}
