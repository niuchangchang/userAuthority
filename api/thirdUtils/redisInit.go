/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-15 14:11:40
 * @LastEditTime: 2019-08-15 14:26:26
 * @LastEditors: Dawn
 */

package thirdUtils

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/wangcong0918/sunrise/log"
	"os"

	"time"
)

func InitClient() (redisClient *redis.Client, err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_IP"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PWD"), // no password set
		DB:       4,                      // use default DB
	})
	_, err = client.Ping().Result()
	if err != nil {
		log.Logger.Info("Code2Session------------->", err)
		return nil, err
	}
	return client, nil
}

func RedisSet(key, value string) {
	redisClient, _ := InitClient()
	err := redisClient.Set(key, value, 720*time.Hour).Err()
	if err != nil {
		log.Logger.Info("RedisSet------------->", err)
		//panic(err)
	}
}
func RedisGet(key string) string {
	redisClient, _ := InitClient()
	key, err := redisClient.Get(key).Result()
	if err != nil {
		log.Logger.Info("RedisGet------------->", err)
		//panic(err)
	}
	return key
}

func RedisDel(key string) {
	redisClient, _ := InitClient()
	err := redisClient.Del(key).Err()
	if err != nil {
		log.Logger.Info("RedisDel------------->", err)
		//panic(err)
	}
}

func RedisDelAndSet(key, value string) {
	redisClient, _ := InitClient()
	err := redisClient.Del(key).Err()
	err1 := redisClient.Set(key, value, 720*time.Hour).Err()
	if err != nil || err1 != nil {
		log.Logger.Info("RedisDel------------->", err)
		//panic(err)
	}
}
