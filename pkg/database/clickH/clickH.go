package clickH

import (
	"fmt"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/config"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func NewClient(cfg *config.Config) (*sqlx.DB, error) {
	conStr := fmt.Sprintf("tcp://%s:%s?debug=%s", cfg.ClickHouse.Host, cfg.ClickHouse.Port, cfg.ClickHouse.Debug)
	connect, err := sqlx.Open(
		"clickhouse", conStr)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("ClickHouse New Client:%s", conStr))
	}
	err = connect.Ping()
	if err != nil {
		return nil, fmt.Errorf("error verifying connection:%w", err)
	}
	return connect, nil
}

