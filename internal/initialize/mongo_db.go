package initialize

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/augustus281/cqrs-pattern/global"
)

const (
	_connectTimeout  = 30 * time.Second
	_maxConnIdleTime = 3 * time.Minute
	_minPoolSize     = 20
	_maxPoolSize     = 300
)

func (s *server) InitMongoDB(ctx context.Context) (*mongo.Client, error) {
	clientOptions := options.Client().
		ApplyURI(global.Config.MongoDB.URI).
		SetConnectTimeout(_connectTimeout).
		SetMaxConnIdleTime(_maxConnIdleTime).
		SetMaxPoolSize(_maxPoolSize)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	global.Logger.Info("connect mongodb successfully!")
	return client, nil
}
