package app

import (
	"context"
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
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	syncProducer, err := messageq.NewProducer([]string{cfg.ClickHouse.Host + ":" + cfg.ClickHouse.Port})
	logger.Check("error creating producer: %v\n", err)

	mq := messageq.Newmq(clickLg, syncProducer, cfg.Kafka.Topic)
	//todo:close?
	go mq.Subscribe()
	Services := services.NewServices(ctx, cacheUsers, mq, repo)

	srv := rpcserver.NewServer(cfg.GRPC.Server)
	var l net.Listener
	go func(l *net.Listener) {
		l, err = srv.Run(Services)
		logger.Check("error running server: %v\n", err)
	}(&l)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()
	err = red.Client.Close()
	logger.Check("error closing cache: %v\n", err)
	err = click.Close()
	logger.Check("error closing log: %v\n", err)
	err = db.Close()
	logger.Check("error closing db: %v\n", err)
	err = l.Close()
	logger.Check("error closing server: %v\n", err)

}
