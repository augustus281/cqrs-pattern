package es

import (
	database "github.com/augustus281/cqrs-pattern/database/sqlc"
)

const _eventsCapacity = 10

type pgEventStore struct {
	eventBus EventsBus
	db       database.DBTX
}

func NewPgEventStore(eventBus EventsBus, db database.DBTX) *pgEventStore {
	return &pgEventStore{
		eventBus: eventBus,
	}
}

// func (p *pgEventStore) handleConcurrency(ctx context.Context, tx database.DBTX)
