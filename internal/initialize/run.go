package initialize

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/augustus281/cqrs-pattern/global"
	v1 "github.com/augustus281/cqrs-pattern/internal/order/delivery/http/v1"
	"github.com/gin-gonic/gin"
)

func (s *server) Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.LoadConfig()
	s.InitLogger()

	postgresConn, err := s.InitDB()
	if err != nil {
		global.Logger.Error("error to init postgresql database", zap.Error(err))
	}
	s.postgresConn = postgresConn

	s.InitRedis(ctx)
	s.InitJeagerTracer()

	mongoDBConn, err := s.InitMongoDB(ctx)
	if err != nil {
		global.Logger.Error("error to init mongoDB", zap.Error(err))
	}
	s.mongoClient = mongoDBConn
	defer mongoDBConn.Disconnect(ctx)

	elasticClient, err := s.InitElasticSearch()
	if err != nil {
		global.Logger.Error("errot to init elastic search", zap.Error(err))
	}
	s.elasticClient = elasticClient

	s.InitEventStoreDB()

	s.RunHealthCheck(ctx)

	orderHandlers := v1.NewOrderHandlers(&gin.RouterGroup{}, s.validate, s.orderService, s.metrics)
	orderHandlers.MapRoutes()

	r := s.InitRouter()
	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)
	r.Run(serverAddr)
}
