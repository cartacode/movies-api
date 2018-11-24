package dbh

import (
	"github.com/VuliTv/go-movie-api/libs/envhelp"
	"github.com/go-redis/redis"
)

var redisConfig = &redis.Options{
	Addr:     envhelp.GetEnv("REDIS_ADDRESS", "localhost:6379"),
	Password: envhelp.GetEnv("REDIS_PASSWORD", ""),
	DB:       envhelp.GetEnvInt("REDIS_DB", 0),
}

// NewRedisConnection --
func NewRedisConnection() (*redis.Client, error) {
	log.Infow("new database handler created",
		"connection_string", redisConfig.Addr,
		"database", redisConfig.DB,
	)

	client := redis.NewClient(redisConfig)
	log.Infow("testing connection")

	res, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	log.Debugw("redis connected", "results", res)
	return client, err
}
