package initialize

import (
	"github.com/augustus281/cqrs-pattern/global"
	"github.com/augustus281/cqrs-pattern/pkg/logger"
)

func (s *server) InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
