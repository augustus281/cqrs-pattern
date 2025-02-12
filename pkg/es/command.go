package es

type Command interface {
	GetAggregateID() string
}

type baseCommand struct {
	AggregateID string `json:"aggregate_id" validate:"required,gte=0"`
}

func NewBaseCommand(aggregateID string) Command {
	return &baseCommand{
		AggregateID: aggregateID,
	}
}

func (c *baseCommand) GetAggregateID() string {
	return c.AggregateID
}
