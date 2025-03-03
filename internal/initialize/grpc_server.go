package initialize

import (
	"net"
	"strconv"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/augustus281/cqrs-pattern/global"
	"github.com/augustus281/cqrs-pattern/pkg/constants"
)

const (
	_maxConnectionIdle = 5
	_gRPCTimeout       = 15
	_maxConnectionAge  = 5
	_gRPCTime          = 10
)

func (s *server) newGRPCServer() (func() error, *grpc.Server, error) {
	listen, err := net.Listen(constants.Tcp, strconv.Itoa(global.Config.GRPC.Port))
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
			nil,
		)),
	)

	return listen.Close, grpcServer, nil
}
