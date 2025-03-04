package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/protobuf/types/known/emptypb"

	orderservice "github.com/augustus281/cqrs-pattern/api"
	"github.com/augustus281/cqrs-pattern/global"
	"github.com/augustus281/cqrs-pattern/internal/metrics"
	v1 "github.com/augustus281/cqrs-pattern/internal/order/commands/v1"
	"github.com/augustus281/cqrs-pattern/internal/order/models"
	"github.com/augustus281/cqrs-pattern/internal/order/queries"
	"github.com/augustus281/cqrs-pattern/internal/order/service"
	grpcerrors "github.com/augustus281/cqrs-pattern/pkg/grpc_errors"
	"github.com/augustus281/cqrs-pattern/pkg/tracing"
	"github.com/augustus281/cqrs-pattern/pkg/utils"
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
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "orderGrpcService.CreateOrder")
	defer span.Finish()
	span.LogFields(log.String("req", req.String()))
	s.metrics.CreateOrderGrpcRequests.Inc()

	aggregateID := uuid.NewString()
	command := v1.NewCreateOrderCommand(
		aggregateID,
		models.ShopItemsFromProto(req.GetShopItems()),
		req.GetAccountEmail(),
		req.GetDeliveryAddress(),
	)
	if err := s.validate.StructCtx(ctx, command); err != nil {
		global.Logger.Error(fmt.Sprintf("(validate) aggregateID: {%s}, err: {%v}", aggregateID, err))
		tracing.TraceErr(span, err)
		return nil, s.errResponse(err)
	}

	if err := s.orderService.Commands.CreateOrder.Handle(ctx, command); err != nil {
		global.Logger.Error(fmt.Sprintf("(CreateOrder.Handle) orderID: {%s}, err: {%v}", aggregateID, err))
		return nil, s.errResponse(err)
	}

	global.Logger.Info(fmt.Sprintf("(created order): orderID: {%s}", aggregateID))
	return &orderservice.CreateOrderResponse{AggregateId: aggregateID}, nil
}

func (s *orderGrpcService) PayOrder(ctx context.Context, req *orderservice.PayOrderRequest) (*orderservice.PayOrderResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "orderGrpcService.PayOrder")
	defer span.Finish()
	span.LogFields(log.String("req", req.String()))
	s.metrics.PayOrderGrpcRequests.Inc()

	payment := models.Payment{
		PaymentID: req.GetPayment().GetId(),
		Timestamp: time.Now(),
	}
	command := v1.NewPayOrderCommand(
		payment,
		req.GetAggregateId(),
	)
	if err := s.validate.StructCtx(ctx, command); err != nil {
		global.Logger.Error(fmt.Sprintf("(validate) err: {%v}", err))
		tracing.TraceErr(span, err)
		return nil, s.errResponse(err)
	}

	if err := s.orderService.Commands.OrderPaid.Handle(ctx, command); err != nil {
		global.Logger.Error(fmt.Sprintf("(OrderPaid.Handle) orderID: {%s}, err: {%v}", req.GetAggregateId(), err))
		return nil, s.errResponse(err)
	}

	global.Logger.Info(fmt.Sprintf("(paid order): orderID: {%s}", req.GetAggregateId()))
	return &orderservice.PayOrderResponse{AggregateId: req.GetAggregateId()}, nil
}

func (s *orderGrpcService) SubmitOrder(ctx context.Context, req *orderservice.SubmitOrderRequest) (*orderservice.SubmitOrderResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "orderGrpcService.SubmitOrder")
	defer span.Finish()
	span.LogFields(log.String("req", req.String()))
	s.metrics.SubmitOrderGrpcRequests.Inc()

	command := v1.NewSubmitOrderCommand(req.GetAggregateId())
	if err := s.validate.StructCtx(ctx, command); err != nil {
		global.Logger.Error(fmt.Sprintf("(validate) err: {%v}", err))
		tracing.TraceErr(span, err)
		return nil, s.errResponse(err)
	}

	if err := s.orderService.Commands.SubmitOrder.Handle(ctx, command); err != nil {
		global.Logger.Error(fmt.Sprintf("(SubmitOrder.Handle) orderID: {%s}, err: {%v}", req.GetAggregateId(), err))
		return nil, s.errResponse(err)
	}

	global.Logger.Info(fmt.Sprintf("(submit order): orderID: {%s}", req.GetAggregateId()))
	return &orderservice.SubmitOrderResponse{AggregateId: req.GetAggregateId()}, nil
}

func (s *orderGrpcService) GetOrderByID(ctx context.Context, req *orderservice.GetOrderByIDRequest) (*orderservice.GetOrderByIDResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "orderGrpcService.GetOrderByID")
	defer span.Finish()
	span.LogFields(log.String("req", req.String()))
	s.metrics.GetOrderByIdGrpcRequests.Inc()

	query := queries.NewGetOrderByIDQuery(req.GetAggregateId())
	if err := s.validate.StructCtx(ctx, query); err != nil {
		global.Logger.Error(fmt.Sprintf("(validate) err: {%v}", err))
		tracing.TraceErr(span, err)
		return nil, s.errResponse(err)
	}

	orderProjection, err := s.orderService.Queries.GetOrderByID.Handle(ctx, query)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("(GetOrderByID.Handle) orderID: {%s}, err: {%v}", req.GetAggregateId(), err))
		return nil, s.errResponse(err)
	}

	global.Logger.Info(fmt.Sprintf("(GetOrderByID) AggregateID: {%s}", req.GetAggregateId()))
	global.Logger.Debug(fmt.Sprintf("(GetOrderByID) orderProjection: {%s}", orderProjection.String()))

	return &orderservice.GetOrderByIDResponse{
		Order: models.OrderProjectionToProto(orderProjection),
	}, nil
}

func (s *orderGrpcService) UpdateShoppingCart(ctx context.Context, req *orderservice.UpdateShoppingCartRequest) (*emptypb.Empty, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "orderGrpcService.UpdateShoppingCart")
	defer span.Finish()
	span.LogFields(log.String("req", req.String()))
	s.metrics.UpdateOrderGrpcRequests.Inc()

	command := v1.NewUpdateShoppingCartCommand(req.GetAggregateId(), models.ShopItemsFromProto(req.GetShopItems()))
	if err := s.validate.StructCtx(ctx, command); err != nil {
		global.Logger.Error(fmt.Sprintf("(validate) err: {%v}", err))
		tracing.TraceErr(span, err)
		return nil, s.errResponse(err)
	}

	if err := s.orderService.Commands.UpdateOrder.Handle(ctx, command); err != nil {
		global.Logger.Error(fmt.Sprintf("(UpdateShoppingCart.Handle) orderID: {%s}, err: {%v}", req.GetAggregateId(), err))
		return nil, s.errResponse(err)
	}

	global.Logger.Info(fmt.Sprintf("(UpdateShoppingCart ) AggregateID: {%s}", req.GetAggregateId()))
	return &emptypb.Empty{}, nil
}

func (s *orderGrpcService) CancelOrder(ctx context.Context, req *orderservice.CancelOrderRequest) (*emptypb.Empty, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "orderGrpcService.CancelOrder")
	defer span.Finish()
	span.LogFields(log.String("req", req.String()))
	s.metrics.CancelOrderGrpcRequests.Inc()

	command := v1.NewCancelOrderCommand(req.GetAggregateId(), req.GetCancelReason())
	if err := s.validate.StructCtx(ctx, command); err != nil {
		global.Logger.Error(fmt.Sprintf("(validate) err: {%v}", err))
		tracing.TraceErr(span, err)
		return nil, s.errResponse(err)
	}

	if err := s.orderService.Commands.CancelOrder.Handle(ctx, command); err != nil {
		global.Logger.Error(fmt.Sprintf("(CancelOrder.Handle) orderID: {%s}, err: {%v}", req.GetAggregateId(), err))
		return nil, s.errResponse(err)
	}

	global.Logger.Info(fmt.Sprintf("(CancelOrder ) AggregateID: {%s}", req.GetAggregateId()))
	return &emptypb.Empty{}, nil
}

func (s *orderGrpcService) CompleteOrder(ctx context.Context, req *orderservice.CompleteOrderRequest) (*emptypb.Empty, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "orderGrpcService.CompleteOrder")
	defer span.Finish()
	span.LogFields(log.String("req", req.String()))
	s.metrics.CompleteOrderGrpcRequests.Inc()

	command := v1.NewCompleteOrderCommand(req.GetAggregateId(), time.Now())
	if err := s.validate.StructCtx(ctx, command); err != nil {
		global.Logger.Error(fmt.Sprintf("(validate) err: {%v}", err))
		tracing.TraceErr(span, err)
		return nil, s.errResponse(err)
	}

	if err := s.orderService.Commands.CompleteOrder.Handle(ctx, command); err != nil {
		global.Logger.Error(fmt.Sprintf("(CompleteOrder.Handle) orderID: {%s}, err: {%v}", req.GetAggregateId(), err))
		return nil, s.errResponse(err)
	}

	global.Logger.Info(fmt.Sprintf("(CompleteOrder) AggregateID: {%s}", req.GetAggregateId()))
	return &emptypb.Empty{}, nil
}

func (s *orderGrpcService) ChangeDeliveryAddress(ctx context.Context, req *orderservice.ChangeDeliveryAddressRequest) (*emptypb.Empty, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "orderGrpcService.ChangeDeliveryAddress")
	defer span.Finish()
	span.LogFields(log.String("req", req.String()))
	s.metrics.ChangeAddressOrderGrpcRequests.Inc()

	command := v1.NewChangeDeliveryAddressCommand(req.GetAggregateId(), req.GetDeliveryAddress())
	if err := s.validate.StructCtx(ctx, command); err != nil {
		global.Logger.Error(fmt.Sprintf("(validate) err: {%v}", err))
		tracing.TraceErr(span, err)
		return nil, s.errResponse(err)
	}

	if err := s.orderService.Commands.ChangeDeliveryAddress.Handle(ctx, command); err != nil {
		global.Logger.Error(fmt.Sprintf("(ChangeOrderDeliveryAddress.Handle) orderID: {%s}, err: {%v}", req.GetAggregateId(), err))
		return nil, s.errResponse(err)
	}

	global.Logger.Info(fmt.Sprintf("(ChangeDeliveryAddress ) AggregateID: {%s}", req.GetAggregateId()))
	return &emptypb.Empty{}, nil
}

func (s *orderGrpcService) Search(ctx context.Context, req *orderservice.SearchRequest) (*orderservice.SearchResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "orderGrpcService.Search")
	defer span.Finish()
	span.LogFields(
		log.String("SearchText", req.GetSearchText()),
		log.Int64("page", req.GetPage()),
		log.Int64("size", req.GetSize()),
	)
	s.metrics.SearchOrderGrpcRequests.Inc()

	query := queries.NewSearchOrdersQuery(req.GetSearchText(), utils.NewPaginationQuery(int(req.GetSize()), int(req.GetPage())))
	if err := s.validate.StructCtx(ctx, query); err != nil {
		global.Logger.Error(fmt.Sprintf("(validate) err: {%v}", err))
		tracing.TraceErr(span, err)
		return nil, s.errResponse(err)
	}

	searchResult, err := s.orderService.Queries.SearchOrders.Handle(ctx, query)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("(SearchOrders.Handle) text: {%s}, err: {%v}", req.GetSearchText(), err))
		return nil, s.errResponse(err)
	}

	global.Logger.Info(fmt.Sprintf("(Search result): searchText: {%s}, pagination: {%+v}", req.GetSearchText(), searchResult.Pagination))
	return nil, nil
}

func (s *orderGrpcService) errResponse(err error) error {
	return grpcerrors.ErrResponse(err)
}
