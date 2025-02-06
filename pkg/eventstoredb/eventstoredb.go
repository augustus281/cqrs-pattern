package eventstoredb

import (
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"go.uber.org/zap"

	"github.com/augustus281/cqrs-pattern/global"
	"github.com/augustus281/cqrs-pattern/pkg/config"
)

func NewEventStoreDB(cfg *config.EventStoreConfig) (*esdb.Client, error) {
	settings, err := esdb.ParseConnectionString(cfg.ConnectionString)
	if err != nil {
		global.Logger.Error("error to parse connect string", zap.Error(err))
		return nil, err
	}
	return esdb.NewClient(settings)
}
