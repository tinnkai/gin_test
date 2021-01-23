package gredis

import (
	"encoding/json"
	"gin_test/pkg/setting"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	redisPool *redis.Pool
)

// Setup Initialize the Redis instance
func Setup() {
	// 默认数据库
	defaultSelectDb := setting.RedisSetting.Db

	redisPool = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					_ = c.Close()
					return c, err
				}
			}
			// 选择db
			c.Do("SELECT", defaultSelectDb)
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	// 清理钩子
	cleanupHook()
}

// 清理钩子
func cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		redisPool.Close()
		os.Exit(0)
	}()
}

// Set a key/value
func Set(key string, data interface{}, expire int) error {
	conn := redisPool.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	// 大于0才会启用过期时间
	if expire > 0 {
		_, err = conn.Do("EXPIRE", key, expire)
		if err != nil {
			return err
		}
	}

	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := redisPool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func Get(key string) ([]byte, error) {
	conn := redisPool.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Get get a keys
func GetKeys(key string) ([]string, error) {
	conn := redisPool.Get()
	defer conn.Close()

	iter := 0
	keys := []string{}
	for {
		arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", key))
		if err != nil {
			return keys, err
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if iter == 0 {
			break
		}
	}

	return keys, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := redisPool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
func LikeDeletes(key string) error {
	conn := redisPool.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}

// incr
func Incr(key string, expire int) (int, error) {
	conn := redisPool.Get()
	defer conn.Close()

	val, err := redis.Int(conn.Do("INCR", key))
	if err != nil {
		return val, err
	}

	// 大于0才会启用过期时间
	if expire > 0 {
		_, err = conn.Do("EXPIRE", key, expire)
		if err != nil {
			return 0, err
		}
	}

	return val, nil
}
