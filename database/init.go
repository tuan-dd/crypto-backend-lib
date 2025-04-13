package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/tuan-dd/common-lib/response"
	"github.com/tuan-dd/common-lib/settings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

func connect(dsn string, cfg *settings.SQLSetting) (*Connection, *response.AppError) {
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal(err)
	}
	poolConfig.MaxConns = 20
	poolConfig.MinConns = 0
	poolConfig.MaxConnLifetime = time.Minute * 2

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatal(err)
	}
	db := stdlib.OpenDBFromPool(pool)

	return &Connection{
		RDBMS: cfg.RDBMS,
		db:    db,
		cfg:   cfg,
	}, nil
}

func (c *Connection) DB() *sql.DB {
	return c.db
}

func (c *Connection) Close() error {
	c.db.Close()
	return nil
}

func (c *Connection) HealthCheck(ctx context.Context) *response.AppError {
	if err := c.db.PingContext(ctx); err != nil {
		return response.DatabaseError(fmt.Errorf("%w: %v", ErrHealthCheckFailed, err))
	}
	return nil
}
