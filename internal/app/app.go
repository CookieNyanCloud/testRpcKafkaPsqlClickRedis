package app

import (
	"context"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/cache/redis"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/config"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/repo/clickLog"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/repo/psql"
	"github.com/cookienyancloud/testrpckafkapsqlclick/pkg/cache"
	"github.com/cookienyancloud/testrpckafkapsqlclick/pkg/database/clickH"
	"github.com/cookienyancloud/testrpckafkapsqlclick/pkg/database/postgres"
	lg "github.com/cookienyancloud/testrpckafkapsqlclick/pkg/logger/logger"
	rpcserver "github.com/cookienyancloud/testrpckafkapsqlclick/pkg/server"
)

func Run(configPath string) {
	ctx := context.Background()

	cfg, err := config.Init(configPath)
	if err != nil {
		lg.Errorf("error init config: %v\n", err)
		return
	}

	db, err := postgres.NewClient(cfg.Postgres)
	if err != nil {
		lg.Errorf("error init db: %v\n", err)
		return
	}
	repo := psql.NewUsersRepo(db)

	red, err := cache.NewRedisClient(cfg.Redis.Addr, ctx)
	if err != nil {
		lg.Errorf("error init redis: %v\n", err)
	}
	cache := redis.NewCache(ctx, red)

	click, err := clickH.NewClient(cfg)
	if err != nil {
		lg.Errorf("error init clickHouse: %v\n", err)
	}
	clickLg := clickLog.NewClickHouseLog(click)

	srv := rpcserver.NewServer(lg, cfg.GRPC.Server)

	defer red.Client.Close()
	defer click.Close()
	defer db.Close()

	if err = srv.Run(); err != nil {
		lg.Errorf("error running server: %v\n", err)
		return
	}
}
