package initialize

import (
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/olivere/elastic/v7"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/augustus281/cqrs-pattern/internal/metrics"
	"github.com/augustus281/cqrs-pattern/internal/order/service"
)

type server struct {
	probeServer   *http.Server
	mongoClient   *mongo.Client
	elasticClient *elastic.Client
	postgresConn  *sql.DB
	pgxConn       *pgxpool.Pool
	metrics       *metrics.ESMicroserviceMetrics
	validate      *validator.Validate
	orderService  *service.OrderService
	kafkaConn     *kafka.Conn
	doneCh        chan struct{}
}

func NewServer() *server {
	return &server{
		validate: validator.New(),
		doneCh:   make(chan struct{}),
	}
}
