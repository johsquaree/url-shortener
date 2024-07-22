package storage

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Storage interface {
	Save(shortURL, originalURL string)
	Get(shortURL string) (string, bool)
}

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(addr, password string, db int) *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisStorage{client: client}
}

func (r *RedisStorage) Save(shortURL, originalURL string) {
	r.client.Set(ctx, shortURL, originalURL, 0).Err()
}

func (r *RedisStorage) Get(shortURL string) (string, bool) {
	val, err := r.client.Get(ctx, shortURL).Result()
	if err == redis.Nil {
		return "", false
	} else if err != nil {
		panic(err)
	}
	return val, true
}
