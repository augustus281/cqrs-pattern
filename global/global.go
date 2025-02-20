package global

import (
	"github.com/redis/go-redis/v9"

	database "github.com/augustus281/cqrs-pattern/database/sqlc"
	"github.com/augustus281/cqrs-pattern/pkg/config"
	"github.com/augustus281/cqrs-pattern/pkg/logger"
)

var (
	Config config.Config
	Logger *logger.LoggerZap
	Db     *database.Store
	Rdb    *redis.Client
)
