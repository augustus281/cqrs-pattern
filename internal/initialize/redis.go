package initialize

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/augustus281/cqrs-pattern/global"
)

func (s *server) InitRedis(ctx context.Context) {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization error: %v", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("Connecting redis successfully!")
	global.Rdb = rdb
}
