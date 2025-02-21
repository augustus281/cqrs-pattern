package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/augustus281/cqrs-pattern/internal/metrics"
	"github.com/augustus281/cqrs-pattern/internal/order/service"
)

type orderHandlers struct {
	group        *gin.RouterGroup
	validate     *validator.Validate
	orderService *service.OrderService
	metrics      *metrics.ESMicroserviceMetrics
}

func NewOrderHandlers(
	group *gin.RouterGroup,
	validate *validator.Validate,
	orderService *service.OrderService,
	metrics *metrics.ESMicroserviceMetrics,
) *orderHandlers {
	return &orderHandlers{
		group:        group,
		validate:     validate,
		orderService: orderService,
		metrics:      metrics,
	}
}
