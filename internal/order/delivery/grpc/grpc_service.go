package grpc

import (
	"context"

	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/types/known/emptypb"

	orderservice "github.com/augustus281/cqrs-pattern/api"
	"github.com/augustus281/cqrs-pattern/internal/metrics"
	"github.com/augustus281/cqrs-pattern/internal/order/service"
)

type orderGrpcService struct {
	orderService *service.OrderService
	validate     *validator.Validate
	metrics      *metrics.ESMicroserviceMetrics
}

func NewOrderGrpcService(orderService *service.OrderService, validate *validator.Validate, metrics *metrics.ESMicroserviceMetrics) *orderGrpcService {
	return &orderGrpcService{
		orderService: orderService,
		validate:     validate,
		metrics:      metrics,
	}
}

func (s *orderGrpcService) CreateOrder(ctx context.Context, req *orderservice.CreateOrderRequest) (*orderservice.CreateOrderResponse, error) {
	return nil, nil
}

func (s *orderGrpcService) PayOrder(ctx context.Context, req *orderservice.PayOrderRequest) (*orderservice.PayOrderResponse, error) {
	return nil, nil
}

func (s *orderGrpcService) SubmitOrder(ctx context.Context, req *orderservice.SubmitOrderRequest) (*orderservice.SubmitOrderResponse, error) {
	return nil, nil
}

func (s *orderGrpcService) GetOrderByID(ctx context.Context, req *orderservice.GetOrderByIDRequest) (*orderservice.GetOrderByIDResponse, error) {
	return nil, nil
}

func (s *orderGrpcService) UpdateShoppingCart(ctx context.Context, req *orderservice.UpdateShoppingCartRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *orderGrpcService) CancelOrder(ctx context.Context, req *orderservice.CancelOrderRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *orderGrpcService) CompleteOrder(ctx context.Context, req *orderservice.CompleteOrderRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *orderGrpcService) Search(ctx context.Context, req *orderservice.SearchRequest) (*orderservice.SearchResponse, error) {
	return nil, nil
}

func (s *orderGrpcService) ChangeDeliveryAddress(ctx context.Context, req *orderservice.ChangeDeliveryAddressRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
