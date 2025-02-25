package es

import "context"

type EventsBus interface {
	ProcessEvents(ctx context.Context, events []Event) error
}
