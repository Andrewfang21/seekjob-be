package redis

import (
	"fmt"
	"log"
	"seekjob/config"
	"time"

	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
)

// Handler defines redis handler operations
type Handler interface {
	Get(key string, variable interface{}) (bool, error)
	SetWithExpiry(key string, value interface{}, expiryInSeconds int) error
	Delete(key string) error
}

type handler struct {
	redisClient *redis.Client
}

var singletonHandler handler

func init() {
	fmt.Println("Preparing redis")
	redisCfg := config.Config.RedisCfg
	redisClient := redis.NewClient(&redis.Options{
		Addr:        redisCfg.Address,
		Password:    redisCfg.Password,
		DB:          redisCfg.Database,
		MaxRetries:  redisCfg.MaxRetries,
		DialTimeout: time.Duration(redisCfg.DialTimeoutInSeconds) * time.Second,
	})
	fmt.Println("Redis Connection is up")

	if _, err := redisClient.Ping().Result(); err != nil {
		log.Fatalf("[ERROR] Fatal error ping redis: %s", err)
		return
	}
	singletonHandler = handler{redisClient: redisClient}
}

// GetHandler returns redis handler
func GetHandler() Handler {
	return &singletonHandler
}

func (h *handler) Get(key string, variable interface{}) (bool, error) {
	bytes, err := h.redisClient.Get(key).Bytes()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	err = msgpack.Unmarshal(bytes, &variable)
	return err == nil, err
}

func (h *handler) SetWithExpiry(key string, value interface{}, expiryInSeconds int) error {
	bytes, err := msgpack.Marshal(value)
	if err != nil {
		return err
	}
	return h.redisClient.Set(key, bytes, time.Duration(expiryInSeconds)*time.Second).Err()
}

func (h *handler) Delete(key string) error {
	return h.redisClient.Del(key).Err()
}
