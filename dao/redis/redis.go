package redis

import (
	"bluebell/settings"
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/spf13/viper"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			// viper.GetString("redis.host"),
			// viper.GetInt("redis.port"),
			cfg.Host,
			cfg.Port,
		),
		Password:/*viper.GetString("redis.password"),*/
		cfg.Password,
		DB:/*viper.GetInt("redis.db"),*/
		cfg.DB,
		PoolSize:/*viper.GetInt("redis.pool_size"),*/
		cfg.PoolSize,
	})
	_, err = rdb.Ping().Result()
	return
}
func Close() {
	_ = rdb.Close()
}
