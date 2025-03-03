package initialize

import (
	"fmt"
	"net"
	"strings"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	orderService "github.com/augustus281/cqrs-pattern/api"
	"github.com/augustus281/cqrs-pattern/global"
	grpc2 "github.com/augustus281/cqrs-pattern/internal/order/delivery/grpc"
)

const (
	_maxConnectionIdle = 5
	_gRPCTimeout       = 15
	_maxConnectionAge  = 5
	_gRPCTime          = 10
)

func (s *server) newGRPCServer() (func(), *grpc.Server, error) {
	listen, err := net.Listen("tcp", ":8088")
	if err != nil {
		return nil, nil, errors.Wrapf(err, "net.Listen")
	}

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: _maxConnectionIdle * time.Minute,
			Timeout:           _gRPCTimeout * time.Second,
			MaxConnectionAge:  _maxConnectionAge * time.Minute,
			Time:              _gRPCTime * time.Minute,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
			s.interceptor.Logger,
		)),
	)

	grpcService := grpc2.NewOrderGrpcService(s.orderService, s.validate, s.metrics)
	orderService.RegisterOrderServiceServer(grpcServer, grpcService)
	grpc_prometheus.Register(grpcServer)

	if global.Config.GRPC.Development {
		reflection.Register(grpcServer)
	}

	go func() {
		global.Logger.Info(fmt.Sprintf("%s gRPC server is starting on port: %d", GetMicroserviceName(), global.Config.GRPC.Port))

		if err := grpcServer.Serve(listen); err != nil {
			if errors.Is(err, net.ErrClosed) {
				global.Logger.Warn("⚠️ gRPC server đã đóng.")
			} else {
				global.Logger.Fatal("🚨 gRPC server dừng do lỗi nghiêm trọng", zap.Error(err))
			}
		}
	}()

	shutdown := func() {
		global.Logger.Info("🔄 Đang dừng gRPC server...")
		grpcServer.GracefulStop()
		listen.Close()
		global.Logger.Info("✅ gRPC server đã dừng.")
	}

	return shutdown, grpcServer, nil
}

func GetMicroserviceName() string {
	return fmt.Sprintf("(%s)", strings.ToUpper(global.Config.ServiceName.ServiceName))
}
