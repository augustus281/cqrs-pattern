package initialize

import (
	"github.com/augustus281/cqrs-pattern/global"
	"github.com/gin-gonic/gin"
)

func (s *server) InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	return r
}
