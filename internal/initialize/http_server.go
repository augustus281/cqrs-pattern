package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/augustus281/cqrs-pattern/global"
	v1 "github.com/augustus281/cqrs-pattern/internal/order/delivery/http/v1"
)

func (s *server) runHttpServer() {
	orderHandlers := v1.NewOrderHandlers(&gin.RouterGroup{}, s.validate, s.orderService, s.metrics)
	orderHandlers.MapRoutes()

	r := s.InitRouter()
	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)
	r.Run(serverAddr)
}
