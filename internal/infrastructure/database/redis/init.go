package redis

import (
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
)

func Init() (redis.Conn, error) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	c, err := redis.DialURL(fmt.Sprintf("redis://user:@%s:%s/0", host, port))
	if err != nil {
		return nil, err
	}
	return c, nil

}
