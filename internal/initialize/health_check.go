package initialize

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/heptiolabs/healthcheck"

	"github.com/augustus281/cqrs-pattern/global"
	"github.com/augustus281/cqrs-pattern/pkg/constants"
)

const (
	_readTimeout  = 15 * time.Second
	_writeTimeout = 15 * time.Second
)

func (s *server) RunHealthCheck(ctx context.Context) {
	health := healthcheck.NewHandler()

	mux := http.NewServeMux()
	s.probeServer = &http.Server{
		Handler:      mux,
		Addr:         global.Config.Probes.Port,
		WriteTimeout: _writeTimeout,
		ReadTimeout:  _readTimeout,
	}
	mux.HandleFunc(global.Config.Probes.LivenessPath, health.LiveEndpoint)
	mux.HandleFunc(global.Config.Probes.ReadinessPath, health.ReadyEndpoint)

	s.configureHealthCheck(ctx, health)

	go func() {
		global.Logger.Info(fmt.Sprintf("(%s) Kubernetes probes listening on port: {%s}", global.Config.ServiceName, global.Config.Probes.Port))
		if err := s.probeServer.ListenAndServe(); err != nil {
			global.Logger.Error("(ListenAndServe) err:", zap.Error(err))
		}
	}()
}

func (s *server) configureHealthCheck(ctx context.Context, health healthcheck.Handler) {
	// Healthcheck PostgreSQL Database
	health.AddReadinessCheck(constants.Postgres, healthcheck.AsyncWithContext(ctx, func() error {
		if err := s.postgresConn.Ping(); err != nil {
			global.Logger.Warn("Postgres Readiness Check", zap.Error(err))
			return err
		}
		return nil
	}, time.Duration(global.Config.Probes.CheckIntervalSeconds)*time.Second))

	health.AddLivenessCheck(constants.Postgres, healthcheck.AsyncWithContext(ctx, func() error {
		if err := s.postgresConn.Ping(); err != nil {
			global.Logger.Warn("Postgres Liveness Check", zap.Error(err))
			return err
		}
		return nil
	}, time.Duration(global.Config.Probes.CheckIntervalSeconds)*time.Second))

	// Healthcheck MongoDB
	health.AddReadinessCheck(constants.MongoDB, healthcheck.AsyncWithContext(ctx, func() error {
		if err := s.mongoClient.Ping(ctx, nil); err != nil {
			global.Logger.Warn("MongoDB Readiness Check", zap.Error(err))
			return err
		}
		return nil
	}, time.Duration(global.Config.Probes.CheckIntervalSeconds)*time.Second))

	health.AddLivenessCheck(constants.Postgres, healthcheck.AsyncWithContext(ctx, func() error {
		if err := s.mongoClient.Ping(ctx, nil); err != nil {
			global.Logger.Warn("MongoDB Liveness Check", zap.Error(err))
			return err
		}
		return nil
	}, time.Duration(global.Config.Probes.CheckIntervalSeconds)*time.Second))
}
