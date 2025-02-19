package es

import (
	"context"
	"fmt"
)

const (
	aggregateStartVersion                = -1 // used for EventStoreDB
	aggregateAppliedEventsInitialCap     = 10
	aggregateUncommittedEventsInitialCap = 10
)

type HandleCommand interface {
	HandleCommand(ctx context.Context, command Command) error
}

type When interface {
	When(event Event) error
}

type when func(event Event) error

type Apply interface {
	Apply(event Event) error
}

type Load interface {
	Load(events []Event) error
}

type Aggregate interface {
	When
	AggregateRoot
}

type AggregateRoot interface {
	GetUncommittedEvents() []Event
	GetID() string
	SetID(id string) *AggregateBase
	GetVersion() int64
	ClearUncommittedEvents()
	ToSnapshot()
	SetType(aggregateType AggregateType)
	GetType() AggregateType
	SetAppliedEvents(events []Event)
	GetAppliedEvents() []Event
	RaiseEvent(event Event) error
	String() string
	Load
	Apply
}

type AggregateType string

type AggregateBase struct {
	ID                string
	Version           int64
	AppliedEvents     []Event
	UncommittedEvents []Event
	Type              AggregateType
	withAppliedEvents bool
	when              when
}

func NewAggregateBase(when when) *AggregateBase {
	if when == nil {
		return nil
	}
	return &AggregateBase{
		Version:           aggregateStartVersion,
		AppliedEvents:     make([]Event, 0, aggregateAppliedEventsInitialCap),
		UncommittedEvents: make([]Event, 0, aggregateUncommittedEventsInitialCap),
		when:              when,
		withAppliedEvents: false,
	}
}

func (a *AggregateBase) SetID(id string) *AggregateBase {
	a.ID = fmt.Sprintf("%s-%s", a.GetType(), id)
	return a
}

func (a *AggregateBase) GetID() string {
	return a.ID
}

func (a *AggregateBase) SetType(aggregateType AggregateType) {
	a.Type = aggregateType
}

func (a *AggregateBase) GetType() AggregateType {
	return a.Type
}

func (a *AggregateBase) GetVersion() int64 {
	return a.Version
}

func (a *AggregateBase) ClearUncommittedEvents() {
	a.UncommittedEvents = make([]Event, 0, aggregateUncommittedEventsInitialCap)
}

func (a *AggregateBase) GetAppliedEvents() []Event {
	return a.AppliedEvents
}

func (a *AggregateBase) SetAppliedEvents(events []Event) {
	a.AppliedEvents = events
}

func (a *AggregateBase) GetUncommittedEvents() []Event {
	return a.UncommittedEvents
}

// Load add existing events from event store to aggregate using When interface method
func (a *AggregateBase) Load(events []Event) error {
	for _, evt := range events {
		if evt.GetAggregateID() != a.GetID() {
			return ErrInvalidAggregate
		}

		if err := a.when(evt); err != nil {
			return err
		}

		if a.withAppliedEvents {
			a.AppliedEvents = append(a.AppliedEvents, evt)
		}
		a.Version++
	}
	return nil
}

// Apply push event to aggregate uncommitted events using When method
func (a *AggregateBase) Apply(event Event) error {
	if event.GetAggregateID() != a.GetID() {
		return ErrInvalidAggregateID
	}
	event.SetAggregateType(a.GetType())
	if err := a.when(event); err != nil {
		return err
	}
	a.Version++
	event.SetVersion(a.GetVersion())
	a.UncommittedEvents = append(a.UncommittedEvents, event)
	return nil
}

// RaiseEvent push event to aggregate applied events using When method, used for load directly from eventstore
func (a *AggregateBase) RaiseEvent(event Event) error {
	if event.GetAggregateID() != a.GetID() {
		return ErrInvalidAggregateID
	}
	if a.GetVersion() >= event.GetVersion() {
		return ErrInvalidEventVersion
	}

	event.SetAggregateType(a.GetType())

	if err := a.when(event); err != nil {
		return err
	}

	if a.withAppliedEvents {
		a.AppliedEvents = append(a.AppliedEvents, event)
	}

	a.Version = event.GetVersion()
	return nil
}

// ToSnapshot prepare AggregateBase for saving Snapshot.
func (a *AggregateBase) ToSnapshot() {
	if a.withAppliedEvents {
		a.AppliedEvents = append(a.AppliedEvents, a.UncommittedEvents...)
	}
	a.ClearUncommittedEvents()
}

func (a *AggregateBase) String() string {
	return fmt.Sprintf("ID: {%s}, Version: {%v}, Type: {%v}, AppliedEvents: {%v}, UncommittedEvents: {%v}",
		a.GetID(),
		a.GetVersion(),
		a.GetType(),
		len(a.GetAppliedEvents()),
		len(a.GetUncommittedEvents()),
	)
}
