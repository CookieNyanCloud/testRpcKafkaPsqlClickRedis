package postgres

import (
	"fmt"

	. "github.com/cookienyancloud/testrpckafkapsqlclick/internal/config"
	"github.com/jmoiron/sqlx"
)

func NewClient(cfg PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, fmt.Errorf("error creating PostgreSQL client:%w\nhost: %v\nport: %v\nuser: %v\ndbname: %v\nsslmode: %v",
			err, cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.SSLMode)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error verifying connection:%w", err)
	}
	return db, nil
}
