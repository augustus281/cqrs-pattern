package initialize

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/avast/retry-go"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	database "github.com/augustus281/cqrs-pattern/database/sqlc"
	"github.com/augustus281/cqrs-pattern/global"
)

const (
	maxConn           = 50
	healthCheckPeriod = 1 * time.Minute
	maxConnIdleTime   = 1 * time.Minute
	maxConnLifetime   = 3 * time.Minute
	minConns          = 10
	lazyConnect       = false
)

func (s *server) InitDB() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		"root",
		global.Config.PostgreSQL.Password,
		"localhost",
		5432,
		global.Config.PostgreSQL.DBName,
	)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Failed to connect database", err)
		return nil, err
	}
	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)

	global.Logger.Info("Connect database successfully!")
	global.Db = database.NewStore(conn)
	return conn, nil
}

func (s *server) InitDBV2(ctx context.Context) error {
	retryOptions := []retry.Option{
		retry.Attempts(uint(global.Config.PostgreSQL.InitRetryCount)),
		retry.Delay(time.Duration(global.Config.PostgreSQL.InitMilliseconds) * time.Millisecond),
		retry.DelayType(retry.BackOffDelay),
		retry.LastErrorOnly(true),
		retry.Context(ctx),
		retry.OnRetry(func(n uint, err error) {
			global.Logger.Error(fmt.Sprintf("retry connect postgres err: %v", err))
		}),
	}

	return retry.Do(func() error {
		pgxConn, err := NewPgxConn(ctx)
		if err != nil {
			return errors.Wrap(err, "postgresql.NewPgxConn")
		}
		s.pgxConn = pgxConn
		return nil
	}, retryOptions...)
}

func NewPgxConn(ctx context.Context) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		"root",
		global.Config.PostgreSQL.Password,
		"localhost",
		5432,
		global.Config.PostgreSQL.DBName,
	)
	poolCfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	poolCfg.MaxConns = maxConn
	poolCfg.HealthCheckPeriod = healthCheckPeriod
	poolCfg.MaxConnIdleTime = maxConnIdleTime
	poolCfg.MaxConnLifetime = maxConnLifetime
	poolCfg.MinConns = minConns
	poolCfg.LazyConnect = lazyConnect

	connPool, err := pgxpool.ConnectConfig(ctx, poolCfg)
	if err != nil {
		return nil, errors.Wrap(err, "pgx.ConnectConfig")
	}

	global.Logger.Info("new pgx conn successfully!")
	return connPool, nil
}
