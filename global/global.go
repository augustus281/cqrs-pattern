package global

import (
	database "github.com/augustus281/cqrs-pattern/database/sqlc"
	"github.com/augustus281/cqrs-pattern/pkg/config"
	"github.com/augustus281/cqrs-pattern/pkg/logger"
)

var (
	Config config.Config
	Logger *logger.LoggerZap
	Db     *database.Store
)
