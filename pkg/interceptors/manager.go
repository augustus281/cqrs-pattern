package interceptors

import (
	"context"
	"time"

	"github.com/augustus281/cqrs-pattern/global"
	"github.com/augustus281/cqrs-pattern/pkg/constants"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GRPCMetricsCb func(err error)

type InterceptorManager interface {
	Logger(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error)
	ClientRequestLoggerInterceptor() func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error
}

type interceptorManager struct {
	metricsCb GRPCMetricsCb
}

func NewInterceptorManager(metricsCb GRPCMetricsCb) *interceptorManager {
	return &interceptorManager{
		metricsCb: metricsCb,
	}
}

func (im *interceptorManager) Logger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	start := time.Now()
	md, _ := metadata.FromIncomingContext(ctx)
	reply, err := handler(ctx, req)
	im.GRPCMiddlewareAccessLogger(info.FullMethod, time.Since(start), md, err)
	if im.metricsCb != nil {
		im.metricsCb(err)
	}
	return reply, nil
}

func (im *interceptorManager) ClientRequestLoggerInterceptor() func(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		md, _ := metadata.FromIncomingContext(ctx)
		im.GRPCClietInterceptorLogger(method, req, reply, time.Since(start), md, err)
		return err
	}
}

func (im *interceptorManager) GRPCMiddlewareAccessLogger(method string, time time.Duration, metaData map[string][]string, err error) {
	if err != nil {
		global.Logger.Error(
			constants.GRPC,
			zap.String(constants.METHOD, method),
			zap.Duration(constants.TIME, time),
			zap.Any(constants.METADATA, metaData),
			zap.Error(err),
		)
		return
	}
	global.Logger.Info(
		constants.GRPC,
		zap.String(constants.METHOD, method),
		zap.Duration(constants.TIME, time),
		zap.Any(constants.METADATA, metaData),
	)
}

func (im *interceptorManager) GRPCClietInterceptorLogger(method string, req interface{}, reply interface{}, time time.Duration, metaData map[string][]string, err error) {
	if err != nil {
		global.Logger.Error(
			constants.GRPC,
			zap.String(constants.METHOD, method),
			zap.Any(constants.REQUEST, req),
			zap.Any(constants.REPLY, reply),
			zap.Duration(constants.TIME, time),
			zap.Any(constants.METADATA, metaData),
			zap.Error(err),
		)
		return
	}
	global.Logger.Info(
		constants.GRPC,
		zap.String(constants.METHOD, method),
		zap.Any(constants.REQUEST, req),
		zap.Any(constants.REPLY, reply),
		zap.Duration(constants.TIME, time),
		zap.Any(constants.METADATA, metaData),
	)
}
