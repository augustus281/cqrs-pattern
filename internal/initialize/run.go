package initialize

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/augustus281/cqrs-pattern/global"
)

func (s *server) Run() {
	LoadConfig()
	InitLogger()

	postgresConn, err := InitDB()
	if err != nil {
		global.Logger.Error("error to init postgresql database", zap.Error(err))
	}
	s.postgresConn = postgresConn

	InitJeagerTracer()
	InitEventStoreDB()

	elasticClient, err := InitElasticSearch()
	if err != nil {
		global.Logger.Error("errot to init elastic search", zap.Error(err))
	}
	s.elasticClient = elasticClient

	mongoDBConn, err := InitMongoDB(context.Background())
	if err != nil {
		global.Logger.Error("error to init mongoDB", zap.Error(err))
	}
	s.mongoClient = mongoDBConn
	defer mongoDBConn.Disconnect(context.TODO())

	r := InitRouter()
	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)
	r.Run(serverAddr)
}
