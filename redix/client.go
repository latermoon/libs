package redix

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type Client interface {
	Call(method string, args ...interface{}) (reply interface{}, err error)
}

func NewClient(server string, maxIdle int) Client {
	c := &client{}

	// https://godoc.org/github.com/garyburd/redigo/redis#NewPool
	c.pool = &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			// if _, err := c.Do("AUTH", password); err != nil {
			// 	c.Close()
			// 	return nil, err
			// }
			return c, err
		},
		// TestOnBorrow: func(c redis.Conn, t time.Time) error {
		// 	_, err := c.Do("PING")
		// 	return err
		// },
	}
	return c
}

type client struct {
	pool *redis.Pool
}

func (c *client) Call(method string, args ...interface{}) (reply interface{}, err error) {
	conn := c.pool.Get()
	defer conn.Close()

	return conn.Do(method, args...)
}
