package database

import (
	"database/sql"
	"fmt"

	"github.com/tuan-dd/crypto-backend-lib/response"
	"github.com/tuan-dd/crypto-backend-lib/settings"

	_ "github.com/lib/pq"
)

type Connection struct {
	db    *sql.DB
	cfg   *settings.SQLSetting
	RDBMS string
}

func NewConnection(cfg *settings.SQLSetting) (*Connection, *response.AppError) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBname, cfg.SSLMode,
	)
	return connect(dsn, cfg)
}
