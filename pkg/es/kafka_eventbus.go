package es

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"

	"github.com/augustus281/cqrs-pattern/global"
	kafkaClient "github.com/augustus281/cqrs-pattern/pkg/kafka"
	"github.com/augustus281/cqrs-pattern/pkg/tracing"
)

type KafkaEventsBusConfig struct {
	Topic             string `mapstructure:"topic" validate:"required"`
	TopicPrefix       string `mapstructure:"topicPrefix" validate:"required"`
	Partitions        int    `mapstructure:"partitions" validate:"required,gte=0"`
	ReplicationFactor int    `mapstructure:"replicationFactor" validate:"required,gte=0"`
	Headers           []kafka.Header
}

type kafkaEventsBus struct {
	producer kafkaClient.Producer
	cfg      KafkaEventsBusConfig
}

func GetTopicName(eventStorePrefix, aggregatType string) string {
	return fmt.Sprintf("%s_%s", eventStorePrefix, aggregatType)
}

func NewKafkaEventsBus(producer kafkaClient.Producer, cfg KafkaEventsBusConfig) *kafkaEventsBus {
	return &kafkaEventsBus{
		producer: producer,
		cfg:      cfg,
	}
}

// ProcessEvents serialize to json and publish es.Event's to the kafka topic.
func (e *kafkaEventsBus) ProcessEvents(ctx context.Context, events []Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "kafkaEventBus.ProcessEvents")
	defer span.Finish()

	eventsBytes, err := json.Marshal(events)
	if err != nil {
		return tracing.TraceWithErr(span, err)
	}

	return e.producer.PublishMessage(ctx, kafka.Message{
		Topic:   GetTopicName(global.Config.KafkaPublisher.TopicPrefix, string(events[0].GetAggregateType())),
		Value:   eventsBytes,
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
		Time:    time.Now().UTC(),
	})
}

func GetKafkaAggregateTopic(aggregateType string) kafka.TopicConfig {
	return kafka.TopicConfig{
		Topic:             GetTopicName(global.Config.KafkaPublisher.TopicPrefix, aggregateType),
		NumPartitions:     global.Config.KafkaPublisher.Partitions,
		ReplicationFactor: global.Config.KafkaPublisher.ReplicationFactor,
	}
}
