package kafka

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"

	"github.com/augustus281/cqrs-pattern/pkg/tracing"
)

type Producer interface {
	PublishMessage(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

type producer struct {
	logger  Logger
	brokers []string
	w       *kafka.Writer
}

func NewProducer(logger Logger, brokers []string) *producer {
	return &producer{
		logger:  logger,
		brokers: brokers,
		w:       NewWriter(brokers, kafka.LoggerFunc(logger.Errorf)),
	}
}

func NewAsyncProducer(logger Logger, brokers []string) *producer {
	return &producer{
		logger:  logger,
		brokers: brokers,
		w:       NewAsyncWriter(brokers, kafka.LoggerFunc(logger.Errorf), logger),
	}
}

func NewAsyncProducerWithCallback(log Logger, brokers []string, cb AsyncWritterCallback) *producer {
	return &producer{
		logger:  log,
		brokers: brokers,
		w:       NewAsyncWritterCallback(brokers, kafka.LoggerFunc(log.Errorf), log, cb),
	}
}

func NewRequireNoneProducer(log Logger, brokers []string) *producer {
	return &producer{
		logger:  log,
		brokers: brokers,
		w:       NewRequireNoneWriter(brokers, kafka.LoggerFunc(log.Errorf), log),
	}
}

func (p *producer) PublishMessage(ctx context.Context, msgs ...kafka.Message) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "producer.PublishMessage")
	defer span.Finish()

	if err := p.w.WriteMessages(ctx, msgs...); err != nil {
		tracing.TraceErr(span, err)
		return err
	}
	return nil
}

func (p *producer) Close() error {
	return p.w.Close()
}
