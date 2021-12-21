package clickH

import (
	"fmt"

	"database/sql"

	"github.com/ClickHouse/clickhouse-go"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/config"
	"github.com/pkg/errors"
)

func NewClient(cfg *config.ClickHouseConfig) (*sql.DB, error) {
	conStr := fmt.Sprintf("tcp://%s:%s?debug=%s", cfg.Host, cfg.Port, cfg.Debug)
	connect, err := sql.Open(
		"clickhouse", conStr)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("ClickHouse New Client:%s", conStr))
	}
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return &sql.DB{}, nil
	}
	return connect, nil
}
