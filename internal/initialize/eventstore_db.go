package initialize

import (
	"github.com/augustus281/cqrs-pattern/global"
	"github.com/augustus281/cqrs-pattern/pkg/eventstoredb"
)

func InitEventStoreDB() {
	db, err := eventstoredb.NewEventStoreDB(&global.Config.EventStore)
	if err != nil {
		global.Logger.Error("init event store db failed!")
		return
	}
	global.Logger.Info("init event store db successfully!")
	defer db.Close()
}
