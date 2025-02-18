package initialize

import (
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type server struct {
	probeServer   *http.Server
	mongoClient   *mongo.Client
	elasticClient *elastic.Client
	postgresConn  *sql.DB
	validate      *validator.Validate
	doneCh        chan struct{}
}

func NewServer() *server {
	return &server{
		validate: validator.New(),
		doneCh:   make(chan struct{}),
	}
}
