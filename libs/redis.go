package libs

import (
	"time"

	"github.com/revel/revel"
	"github.com/robfig/config"
	"gopkg.in/redis.v5"
)

// RedisClient rs client
var RedisClient *redis.Client

// InitRedis 初始化
func InitRedis() {
	c, _ := config.ReadDefault(revel.BasePath + "/conf/database.conf")
	addr, _ := c.String(revel.RunMode, "redis.addr")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:         addr,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
}
