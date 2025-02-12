package store

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/augustus281/cqrs-pattern/global"
	"github.com/augustus281/cqrs-pattern/pkg/es"
	"github.com/augustus281/cqrs-pattern/pkg/tracing"
)

type eventStore struct {
	db *esdb.Client
}

func NewEventStore(db *esdb.Client) *eventStore {
	return &eventStore{
		db: db,
	}
}

func (e *eventStore) SaveEvents(ctx context.Context, streamID string, events []es.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "eventStore.SaveEvents")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", streamID))

	eventsData := make([]esdb.EventData, 0, len(events))
	for _, event := range events {
		eventsData = append(eventsData, event.ToEventData())
	}

	stream, err := e.db.AppendToStream(ctx, streamID, esdb.AppendToStreamOptions{}, eventsData...)
	if err != nil {
		tracing.TraceErr(span, err)
		return err
	}

	global.Logger.Debug(fmt.Sprintf("SaveEvents stream: %+v", stream))
	return nil
}

func (e *eventStore) LoadEvents(ctx context.Context, streamID string) ([]es.Event, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "eventStore.SaveEvents")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", streamID))

	stream, err := e.db.ReadStream(ctx, streamID, esdb.ReadStreamOptions{
		Direction: esdb.Forwards,
		From:      esdb.Revision(1),
	}, 100)
	if err != nil {
		tracing.TraceErr(span, err)
		return nil, err
	}
	defer stream.Close()

	events := make([]es.Event, 0, 100)
	for {
		event, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			tracing.TraceErr(span, err)
			return nil, err
		}
		events = append(events, es.NewEventFromRecorded(event.Event))
	}
	return events, nil
}
