package app

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	rpcserver "github.com/cookienyancloud/testrpckafkapsqlclick/pkg/server"
	_ "github.com/lib/pq"
)

func Run(configPath string) {

	ctx := context.Background()

	cfg, err := config.Init(configPath)
	logger.Check("error init config: %v", err)

	db, err := postgres.NewClient(&cfg.Postgres)
	logger.Check("error init db: %v", err)
	repo := psql.NewUsersRepo(db)
	err = repo.Migrate()
	logger.Check("error migrating db: %v", err)
	red, err := cache.NewRedisClient(cfg.Redis.Addr, ctx)
	logger.Check("error init redis: %v", err)

	cacheUsers := redis.NewCache(red)

	click, err := clickH.NewClient(&cfg.ClickHouse)
	logger.Check("error init clickHouse: %v", err)
	clickLg := clickLog.NewClickHouseLog(click)
	err = clickLg.Migrate()
	logger.Check("error migrating logs: %v", err)
	syncProducer, err := messageq.NewProducer([]string{cfg.Kafka.Addr})
	logger.Check("error creating producer: %v", err)
	mq := messageq.Newmq(clickLg, syncProducer, cfg.Kafka.Topic)
	//todo:close?
	go mq.Subscribe()
	Services := services.NewServices(ctx, cacheUsers, mq, repo)

	srv := rpcserver.NewServer(cfg.GRPC.Server)
	var l net.Listener
	go func(l *net.Listener) {
		l, err = srv.Run(Services)
		logger.Check("error running server: %v", err)
	}(&l)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()
	err = red.Client.Close()
	logger.Check("error closing cache: %v", err)
	err = click.Close()
	logger.Check("error closing log: %v", err)
	err = db.Close()
	logger.Check("error closing db: %v", err)
	err = l.Close()
	logger.Check("error closing server: %v", err)
	err = syncProducer.Close()
	logger.Check("error closing kafka prod: %v", err)
}
