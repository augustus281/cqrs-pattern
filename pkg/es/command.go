package es

type Command interface {
	GetAggregateID() string
}

type BaseCommand struct {
	AggregateID string `json:"aggregate_id" validate:"required,gte=0"`
}

func NewBaseCommand(aggregateID string) BaseCommand {
	return BaseCommand{
		AggregateID: aggregateID,
	}
}

func (c *BaseCommand) GetAggregateID() string {
	return c.AggregateID
}
