package services

import (
	"context"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/cache/redis"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/domain"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/messageq"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/repo/psql"
)

type services struct {
	ctx   context.Context
	cache redis.ICache
	mq    messageq.Imq
	data  psql.IUsersRepo
}

func NewServices(
	ctx context.Context,
	cache redis.ICache,
	mq messageq.Imq,
	data psql.IUsersRepo) IServices {
	return &services{
		ctx:   ctx,
		cache: cache,
		mq:    mq,
		data:  data,
	}
}

type IServices interface {
	ISrvCache
	ISrvDb
	IQueue
}

