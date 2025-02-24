package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
)

func NewKafkaReader(kafkeURLs []string, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:                kafkeURLs,
		GroupID:                groupID,
		Topic:                  topic,
		MinBytes:               minBytes,
		MaxBytes:               maxBytes,
		QueueCapacity:          queueCapacity,
		HeartbeatInterval:      heartbeatInterval,
		CommitInterval:         commitInterval,
		PartitionWatchInterval: partitionWatchInterval,
		MaxAttempts:            maxAttempts,
		MaxWait:                3 * time.Second,
		Dialer:                 &kafka.Dialer{Timeout: dialTimeout},
		ReadBackoffMax:         300 * time.Millisecond,
	})
}
