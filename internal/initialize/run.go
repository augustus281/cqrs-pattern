package initialize

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/augustus281/cqrs-pattern/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitDB()
	InitJeagerTracer()
	InitEventStoreDB()
	_, err := InitElasticSearch()
	if err != nil {
		global.Logger.Error("errot to init elastic search", zap.Error(err))
	}

	_, err = InitMongoDB(context.Background())
	if err != nil {
		global.Logger.Error("error to init mongoDB", zap.Error(err))
	}

	r := InitRouter()
	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)
	r.Run(serverAddr)
}
