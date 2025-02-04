package initialize

import (
	"fmt"

	"github.com/augustus281/cqrs-pattern/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitDB()

	r := InitRouter()
	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)
	r.Run(serverAddr)
}
