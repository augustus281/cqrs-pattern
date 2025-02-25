package initialize

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"

	"github.com/augustus281/cqrs-pattern/global"
	"github.com/augustus281/cqrs-pattern/pkg/constants"
	kafkaClient "github.com/augustus281/cqrs-pattern/pkg/kafka"
)

func (s *server) connectKafkaBrokers(ctx context.Context) error {
	kafkaConn, err := kafkaClient.NewKafkaConn(ctx)
	if err != nil {
		return errors.Wrap(err, "kafka.NewKafkaCon")
	}
	s.kafkaConn = kafkaConn

	brokers, err := kafkaConn.Brokers()
	if err != nil {
		return errors.Wrap(err, "kafkaConn.Brokers")
	}

	global.Logger.Info(fmt.Sprintf("(kakfa connected) brokers: %+v", brokers))
	return nil
}

func (s *server) initKafkaTopics(ctx context.Context) {
	controller, err := s.kafkaConn.Controller()
	if err != nil {
		global.Logger.Error("kafkaConn.Controller err: %v", zap.Error(err))
	}

	controllerURI := net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port))
	global.Logger.Info(fmt.Sprintf("(kafka controller uri) controllerURI: %s", controllerURI))

	conn, err := kafka.DialContext(ctx, constants.Tcp, controllerURI)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("initKafkaTopics.DialContext err: %v", err))
	}
	defer conn.Close()

	global.Logger.Info(fmt.Sprintf("(established new kafka controller connection) controllerURI: %s", controllerURI))
}
