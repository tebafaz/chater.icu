package redis

import "github.com/gomodule/redigo/redis"

func Setex(key string, timeSec int, value string) error {
	_, err := redisConn.Do("SETEX", key, timeSec, value)
	return err
}

func GetStr(key string) (string, error) {
	got, err := redis.String(redisConn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return got, nil
}

func GetInt(key string) (int, error) {
	got, err := redis.Int(redisConn.Do("GET", key))
	if err != nil {
		return 0, err
	}

	return got, nil
}

func Del(token string) error {
	_, err := redisConn.Do("DEL", token)
	return err
}
