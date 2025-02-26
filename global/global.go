package global

import (
	"github.com/redis/go-redis/v9"

	"github.com/augustus281/cqrs-pattern/pkg/config"
	"github.com/augustus281/cqrs-pattern/pkg/logger"
)

var (
	Config config.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
)
