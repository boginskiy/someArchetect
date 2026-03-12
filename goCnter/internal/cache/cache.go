package cache

import (
	"context"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

type Casher interface {
	GetCnt(string) (int, error)
	IncrementCnt(string) error
}

type Cashe struct {
	client *redis.Client
}

func NewCashe() *Cashe {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	return &Cashe{client: client}
}

func (c *Cashe) GetCnt(id string) (int, error) {
	v, err := c.client.Get(id).Int()
	if err == redis.Nil {
		// Счетчик отсутствует
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return v, nil
}

func (c *Cashe) IncrementCnt(id string) error {
	_, err := c.client.Incr(id).Result()
	return err
}
