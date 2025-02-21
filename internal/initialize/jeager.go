package initialize

import (
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"github.com/augustus281/cqrs-pattern/global"
	"github.com/augustus281/cqrs-pattern/pkg/tracing"
)

// Approach 1: Using jaeger
func (s *server) InitJeagerTracer() {
	if global.Config.Jaeger.Enable {
		tracer, closer, err := tracing.NewJaegerTracer(&global.Config.Jaeger)
		if err != nil {
			global.Logger.Error("error to new jaeger tracer")
			return
		}
		defer closer.Close()
		opentracing.SetGlobalTracer(tracer)
	}
}

// Approach 2: Using open-telemtry
func StartTracing() {
	_, err := tracing.StartTracing()
	if err != nil {
		global.Logger.Error("error to start tracing", zap.Error(err))
		return
	}
	global.Logger.Info("start tracing successfully!")
}
