package store

import (
	"context"
	"fmt"
	"io"
	"math"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"

	"github.com/augustus281/cqrs-pattern/global"
	"github.com/augustus281/cqrs-pattern/pkg/es"
	"github.com/augustus281/cqrs-pattern/pkg/tracing"
)

const _count = math.MaxInt64

var tracer = otel.Tracer("github.com/augustus281/cqrs-pattern/pkg/es/store")

type aggregateStore struct {
	db *esdb.Client
}

func NewAggregateStore(db *esdb.Client) *aggregateStore {
	return &aggregateStore{
		db: db,
	}
}

func (a *aggregateStore) Load(ctx context.Context, aggregate es.Aggregate) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "aggregateStore.Load")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", aggregate.GetID()))

	stream, err := a.db.ReadStream(ctx, aggregate.GetID(), esdb.ReadStreamOptions{}, _count)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "db.ReadStream")
	}
	defer stream.Close()

	for {
		event, err := stream.Recv()
		if errors.Is(err, esdb.ErrStreamNotFound) {
			tracing.TraceErr(span, err)
			return errors.Wrap(err, "stream.Recv")
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			tracing.TraceErr(span, err)
			return errors.Wrap(err, "stream.Recv")
		}

		esEvent := es.NewEventFromRecorded(event.Event)
		if err := aggregate.RaiseEvent(esEvent); err != nil {
			tracing.TraceErr(span, err)
			return errors.Wrap(err, "RaiseEvent")
		}
		global.Logger.Debug("(Load) esEvent ", zap.String("event", aggregate.String()))
	}
	global.Logger.Debug("(Load) aggregate ", zap.String("aggregate", aggregate.String()))
	return nil
}

func (a *aggregateStore) Save(ctx context.Context, aggregate es.Aggregate) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "aggregateStore.Save")
	defer span.Finish()
	span.LogFields(log.String("aggregate", aggregate.String()))

	if len(aggregate.GetUncommittedEvents()) == 0 {
		global.Logger.Debug("(Save) [no uncommittedEvents] ", zap.Int("len", len(aggregate.GetUncommittedEvents())))
		return nil
	}

	eventsData := make([]esdb.EventData, 0, len(aggregate.GetUncommittedEvents()))
	for _, event := range aggregate.GetUncommittedEvents() {
		eventsData = append(eventsData, event.ToEventData())
	}

	var expectedRevision esdb.ExpectedRevision
	if aggregate.GetVersion() == 0 {
		expectedRevision = esdb.NoStream{}
		global.Logger.Debug(fmt.Sprintf("(Save) expectedRevision: {%T}", expectedRevision))

		appendStream, err := a.db.AppendToStream(
			ctx,
			aggregate.GetID(),
			esdb.AppendToStreamOptions{ExpectedRevision: expectedRevision},
			eventsData...,
		)
		if err != nil {
			tracing.TraceErr(span, err)
			return errors.Wrap(err, "db.AppendToStream")
		}

		global.Logger.Debug(fmt.Sprintf("(Save) stream: {%+v}", appendStream))
		return nil
	}

	readOps := esdb.ReadStreamOptions{Direction: esdb.Backwards, From: esdb.End{}}
	stream, err := a.db.ReadStream(context.Background(), aggregate.GetID(), readOps, 1)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "db.ReadStream")
	}
	defer stream.Close()

	lastEvent, err := stream.Recv()
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "stream.Recv")
	}

	expectedRevision = esdb.Revision(lastEvent.OriginalEvent().EventNumber)
	global.Logger.Debug(fmt.Sprintf("(Save) expectedRevision : {%T}", expectedRevision))

	appendStream, err := a.db.AppendToStream(
		ctx,
		aggregate.GetID(),
		esdb.AppendToStreamOptions{ExpectedRevision: expectedRevision},
		eventsData...,
	)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "db.AppendToStream")
	}

	global.Logger.Debug(fmt.Sprintf("(Save) stream: {%+v}", appendStream))
	aggregate.ClearUncommittedEvents()
	return nil
}

func (a *aggregateStore) Exists(ctx context.Context, streamID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "aggregateStore.Exists")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", streamID))

	readStreamOptions := esdb.ReadStreamOptions{
		Direction: esdb.Backwards,
		From:      esdb.Revision(1),
	}
	stream, err := a.db.ReadStream(ctx, streamID, readStreamOptions, 1)
	if err != nil {
		return errors.Wrap(err, "db.ReadStream")
	}
	defer stream.Close()

	for {
		_, err := stream.Recv()
		if errors.Is(err, esdb.ErrStreamNotFound) {
			tracing.TraceErr(span, err)
			return errors.Wrap(esdb.ErrStreamNotFound, "stream.Recv")
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			tracing.TraceErr(span, err)
			return errors.Wrap(err, "stream.Recv")
		}
	}
	return nil
}
