package repository

import (
	"context"

	"github.com/augustus281/cqrs-pattern/internal/dto"
	"github.com/augustus281/cqrs-pattern/internal/order/models"
	"github.com/augustus281/cqrs-pattern/pkg/utils"
)

type ElasticOrderRepository interface {
	IndexOrder(ctx context.Context, order *models.OrderProjection) error
	GetByID(ctx context.Context, orderID string) (*models.OrderProjection, error)
	UpdateOrder(ctx context.Context, order *models.OrderProjection) error
	Search(ctx context.Context, text string, pq *utils.Pagination) (*dto.OrderSearchResponseDTO, error)
}
