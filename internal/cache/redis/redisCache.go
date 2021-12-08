package redis

import (
	"context"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/domain"
	"github.com/cookienyancloud/testrpckafkapsqlclick/pkg/cache"
	"time"
)

type Cache struct {
	rd *cache.RedisClient
}

func NewCache(rd *cache.RedisClient, ) ICache {
	return &Cache{
		rd: rd,
	}
}

type ICache interface {
	CacheAll(ctx context.Context, all []domain.User) error
	GetAll(ctx context.Context) ([]domain.User, error)
}

func (c Cache) CacheAll(ctx context.Context, all []domain.User) error {
	for _, user := range all {
		err := c.rd.Client.Set(ctx, user.Id, user, time.Minute).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c Cache) GetAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	keys, err := c.rd.Client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}
	var user domain.User
	for _, key := range keys {
		err := c.rd.Client.Get(ctx, key).Scan(user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
