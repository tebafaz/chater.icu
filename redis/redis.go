package redis

import (
	"fmt"
	"log"
	"os"

	//"os"

	"github.com/gomodule/redigo/redis"
)

//Redis a
var redisConn redis.Conn

//InitRedis a
func InitRedis() error {
	var err error
	redisConn, err = redis.Dial("tcp", fmt.Sprintf("%s:%s", os.Getenv("HOSTNAME"), os.Getenv("REDIS_PORT")), redis.DialPassword(os.Getenv("REDIS_PASSWORD")))
	if err != nil {
		return err
	}
	fmt.Println("connection to Redis created")
	return nil
}

//CloseRedis a
func CloseRedis() {
	err := redisConn.Close()
	if err != nil {
		log.Fatal("redis refused to close")
	}
}

func UpdateExpireTime(key string, timeSec int) error {
	_, err := redisConn.Do("EXPIRE", key, timeSec)
	return err
}
