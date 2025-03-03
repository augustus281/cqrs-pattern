package initialize

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/augustus281/cqrs-pattern/global"
	"github.com/augustus281/cqrs-pattern/internal/metrics"
	"github.com/augustus281/cqrs-pattern/pkg/interceptors"
)

func (s *server) Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.LoadConfig()
	s.InitLogger()

	s.InitRedis(ctx)
	s.InitJeagerTracer()
	s.metrics = metrics.NewESMicroserviceMetrics()
	s.interceptor = interceptors.NewInterceptorManager(s.getGrpcMetricsCb())

	if err := s.InitDBV2(ctx); err != nil {
		global.Logger.Error("error to init postgresql database", zap.Error(err))
	}
	defer s.pgxConn.Close()
	if err := s.runMigrate(); err != nil {
		global.Logger.Error("failed to run migration database", zap.Error(err))
	}

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

	go s.runHttpServer()

	shutdown, _, err := s.newGRPCServer()
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	shutdown()
}

func (s *server) getGrpcMetricsCb() func(err error) {
	return func(err error) {
		if err != nil {
			s.metrics.ErrorGrpcRequests.Inc()
		} else {
			s.metrics.SuccessGrpcRequests.Inc()
		}
	}
}
