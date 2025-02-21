package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"

	"github.com/augustus281/cqrs-pattern/global"
	"github.com/augustus281/cqrs-pattern/internal/dto"
	v1 "github.com/augustus281/cqrs-pattern/internal/order/commands/v1"
	httpErrors "github.com/augustus281/cqrs-pattern/pkg/http_errors"
	"github.com/augustus281/cqrs-pattern/pkg/tracing"
)

type OrderHandlers interface {
	CreateOrder() gin.HandlerFunc
	PayOrder() gin.HandlerFunc
	SubmitOrder() gin.HandlerFunc
	UpdateShoppingCart() gin.HandlerFunc

	GetOrderByID() gin.HandlerFunc
	Search() gin.HandlerFunc
}

func (h *orderHandlers) CreateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "orderHandlers.CreateOrder")
		defer span.Finish()
		h.metrics.SubmitOrderHttpRequests.Inc()

		var reqDTO dto.CreateOderRequest
		if err := ctx.ShouldBind(&reqDTO); err != nil {
			global.Logger.Error(fmt.Sprintf("(ShouldBind) err : {%v}", err))
			tracing.TraceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, global.Config.Server.Debug)
		}

		if err := h.validate.StructCtx(ctx.Request.Context(), reqDTO); err != nil {
			global.Logger.Error(fmt.Sprintf("(validate) err: {%v}", err))
			tracing.TraceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, global.Config.Server.Debug)
		}

		id := uuid.NewString()
		command := v1.NewCreateOrderCommand(id, reqDTO.ShopItems, reqDTO.AccountEmail, reqDTO.DeliveryAddress)
		err := h.orderService.Commands.CreateOrder.Handle(ctx, command)
		if err != nil {
			global.Logger.Error(fmt.Sprintf("(CreateOrder.Handle) id: {%s}, err: {%v}", id, err))
			tracing.TraceErr(span, err)
			httpErrors.ErrorCtxResponse(ctx, err, global.Config.Server.Debug)
		}

		global.Logger.Info(fmt.Sprintf("(order created) id: {%s}", id))
		ctx.JSON(http.StatusCreated, id)
	}
}

func (h *orderHandlers) PayOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (h *orderHandlers) CancelOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (h *orderHandlers) CompleteOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (h *orderHandlers) ChangeDeliveryAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (h *orderHandlers) SubmitOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func (h *orderHandlers) UpdateShoppingCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (h *orderHandlers) GetOrderByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (h *orderHandlers) Search() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
