package kafka

import (
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
)

func NewWriter(brokers []string, errLogger kafka.Logger) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		MaxAttempts:  writerMaxAttempts,
		ErrorLogger:  errLogger,
		Compression:  compress.Snappy,
		ReadTimeout:  writerReadTimeout,
		WriteTimeout: writerWriteTimeout,
		BatchTimeout: batchTimeout,
		BatchSize:    batchSize,
		Async:        false,
	}
}

func NewAsyncWriter(brokers []string, errLogger kafka.Logger, log Logger) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		MaxAttempts:  writerMaxAttempts,
		ErrorLogger:  errLogger,
		Compression:  compress.Snappy,
		ReadTimeout:  writerReadTimeout,
		WriteTimeout: writerWriteTimeout,
		Async:        true,
		Completion: func(messages []kafka.Message, err error) {
			if err != nil {
				log.Errorf("(kafka.AsyncWriter Error) topic: %s, partition: %v, offset: %v err: %v", messages[0].Topic, messages[0].Partition, messages[0].Offset, err)
				return
			}
		},
	}
}

type AsyncWritterCallback func(message []kafka.Message) error

func NewAsyncWritterCallback(brokers []string, errLogger kafka.Logger, log Logger, callback AsyncWritterCallback) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		MaxAttempts:  writerMaxAttempts,
		ErrorLogger:  errLogger,
		Compression:  compress.Snappy,
		ReadTimeout:  writerReadTimeout,
		WriteTimeout: writerWriteTimeout,
		Async:        true,
		Completion: func(messages []kafka.Message, err error) {
			if err != nil {
				log.Errorf("(kafka.AsyncWriter Error) topic: %s, partition: %v, offset: %v err: %v", messages[0].Topic, messages[0].Partition, messages[0].Offset, err)
				if err := callback(messages); err != nil {
					log.Errorf("(kafka.AsyncWriter Callback Error) err: %v", err)
					return
				}
				return
			}
		},
	}
}

func NewRequireNoneWriter(brokers []string, errLogger kafka.Logger, log Logger) *kafka.Writer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireNone,
		MaxAttempts:  writerMaxAttempts,
		ErrorLogger:  errLogger,
		Compression:  compress.Snappy,
		ReadTimeout:  writerRequireNoneReadTimeout,
		WriteTimeout: writerRequireNoneWriteTimeout,
		Async:        false,
		Completion: func(messages []kafka.Message, err error) {
			if err != nil {
				log.Errorf("(kafka.Writer Error) topic: %s, partition: %v, offset: %v err: %v", messages[0].Topic, messages[0].Partition, messages[0].Offset, err)
				return
			}
		},
	}
	return w
}
