package redis

import (
	"github.com/go-redis/redis"
	"taskService/config"
)

type Redis struct {
	DbConnect *redis.Client
}

func New(db *config.Config) (*Redis, error) {
	client := &Redis{
		DbConnect: redis.NewClient(&redis.Options{
			Addr:     db.Redis.Address,
			Password: db.Redis.Password,
			DB:       db.Redis.Base,
		})}
	_, err := client.DbConnect.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil

}
