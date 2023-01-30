package db

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopartsrv/public/config"
	"time"
)

var (
	RedisPoolMap map[string]*redis.Pool
	redisCfg     map[string]interface{}
)
func init() {
	RedisPoolMap = make(map[string]*redis.Pool)
	setRedisCon()
}

func NewRedisPool(instance string) {
	instanceCfg := redisCfg[instance].(map[string]interface{})
	host := instanceCfg["host"].(string)
	port := instanceCfg["port"].(string)
	redisURL := host + ":" + port
	redisMaxIdle := int(instanceCfg["max_idle"].(int64))
	redisIdleTimeout := time.Duration(instanceCfg["idle_timeout"].(int64))
	redisConnectTimeout := instanceCfg["connect_timeout"].(int64)
	redisReadTimeout := instanceCfg["read_timeout"].(int64)
	redisWriteTimeout := instanceCfg["write_timeout"].(int64)
	RedisPoolMap[instance] = &redis.Pool{
		MaxIdle:     redisMaxIdle,
		IdleTimeout: redisIdleTimeout * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisURL,
				redis.DialConnectTimeout(time.Duration(redisConnectTimeout)*time.Second),
				redis.DialReadTimeout(time.Duration(redisReadTimeout)*time.Second),
				redis.DialWriteTimeout(time.Duration(redisWriteTimeout)*time.Second))
			if err != nil {
				return nil, err
			}
			//if _, authErr := c.Do("AUTH", ""); authErr != nil {
			//	return nil, fmt.Errorf("redis auth password error: %s", authErr)
			//}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}

func setRedisCon() {
	redisCfg = config.GetRedisConfig()
	for k := range redisCfg {
		NewRedisPool(k)
	}
}