package dbh

import (
	"fmt"

	"github.com/VuliTv/go-movie-api/libs/envhelp"
	"github.com/go-redis/redis"
)

// RedisHandler --
type RedisHandler struct {
	*redis.Client
}

var redisConfig = &redis.Options{
	Addr:     fmt.Sprintf("%s:%s", envhelp.GetEnv("REDIS_ADDRESS", "localhost"), envhelp.GetEnv("REDIS_PORT", "6379")),
	Password: envhelp.GetEnv("REDIS_PASSWORD", ""),
	DB:       envhelp.GetEnvInt("REDIS_DB", 0),
}

// New --
func (r *RedisHandler) New(controller string) error {
	log.Infow("new redis handler created",
		"caller", controller,
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

	r.Client = client
	return err
}
