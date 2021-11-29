package redis

import (
	"context"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/domain"
	"github.com/cookienyancloud/testrpckafkapsqlclick/pkg/cache"
	"time"
)

type Cache struct {
	ctx context.Context
	rd  *cache.RedisClient
}

func NewCache(ctx context.Context, rd *cache.RedisClient, ) ICache {
	return &Cache{
		ctx: ctx,
		rd:  rd,
	}
}

type ICache interface {
	CacheAll(all []domain.User) error
	GetAll() ([]domain.User, error)
}

func (c Cache) CacheAll(all []domain.User) error {
	return c.rd.Client.Set(c.ctx, "users", all, time.Minute).Err()
}

func (c Cache) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := c.rd.Client.Get(c.ctx, "users").Scan(&users)
	return users, err
}
