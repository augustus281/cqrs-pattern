package kafka

import (
	"context"

	"github.com/augustus281/cqrs-pattern/global"
	"github.com/segmentio/kafka-go"
)

func NewKafkaConn(ctx context.Context) (*kafka.Conn, error) {
	return kafka.DialContext(ctx, "tcp", global.Config.Kafka.Brokers[0])
}
