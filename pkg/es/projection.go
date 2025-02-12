package es

import "context"

type Projection interface {
	When(ctx context.Context, evt Event) error
}
