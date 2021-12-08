package app

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/cache/redis"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/config"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/messageq"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/repo/clickLog"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/repo/psql"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/services"
	"github.com/cookienyancloud/testrpckafkapsqlclick/pkg/cache"
	"github.com/cookienyancloud/testrpckafkapsqlclick/pkg/database/clickH"
	"github.com/cookienyancloud/testrpckafkapsqlclick/pkg/database/postgres"
	"github.com/cookienyancloud/testrpckafkapsqlclick/pkg/logger"
	lg "github.com/cookienyancloud/testrpckafkapsqlclick/pkg/logger/logger"
	rpcserver "github.com/cookienyancloud/testrpckafkapsqlclick/pkg/server"
)

func Run(configPath string) {

	ctx := context.Background()

	cfg, err := config.Init(configPath)
	logger.Check("error init config: %v\n", err)

	db, err := postgres.NewClient(cfg.Postgres)
	logger.Check("error init db: %v\n", err)
	repo := psql.NewUsersRepo(db)

	red, err := cache.NewRedisClient(cfg.Redis.Addr, ctx)
	logger.Check("error init redis: %v\n", err)

	cacheUsers := redis.NewCache(red)

	click, err := clickH.NewClient(cfg)
	logger.Check("error init clickHouse: %v\n", err)
	clickLg := clickLog.NewClickHouseLog(click)

	syncProducer, err := messageq.NewProducer([]string{cfg.ClickHouse.Host+":"+cfg.ClickHouse.Port})
	logger.Check("error creating producer: %v\n", err)


	mq := messageq.Newmq(clickLg, syncProducer, cfg.Kafka.Topic)


	Services := services.NewServices(cacheUsers, repo)

	srv := rpcserver.NewServer(cfg.GRPC.Server)

	defer red.Client.Close()
	defer click.Close()
	defer db.Close()

	if err = srv.Run(Services); err != nil {
		lg.Errorf("error running server: %v\n", err)
		return
	}
}
