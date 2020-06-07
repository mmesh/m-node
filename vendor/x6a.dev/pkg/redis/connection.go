// Copyright (C) 2019 <x6a@7n.io>
//
// pkg is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// pkg is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with pkg. If not, see <http://www.gnu.org/licenses/>.

package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"x6a.dev/pkg/errors"
)

var RedisURL string
var Pool *redis.Pool
var Close = make(chan struct{})

func init() {
	Pool = newPool()

	go func() {
		<-Close
		Pool.Close()
	}()
}

func redisDial(RedisURL string) (redis.Conn, error) {
	return redis.DialURL(
		RedisURL,
		redis.DialConnectTimeout(10*time.Second),
		redis.DialReadTimeout(5*time.Second),
		redis.DialKeepAlive(3*time.Minute), // default: 5 minutes
	)
}

func redisConnect() (redis.Conn, error) {
	if len(RedisURL) == 0 {
		return nil, errors.New("redis config not set")
	}

	c, err := redisDial(RedisURL)
	for i := 0; err != nil && i < 10; i++ {
		fmt.Println("WARNING: unable to connect to redis, retrying in 3s..")
		time.Sleep(3 * time.Second)
		c, err = redisDial(RedisURL)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] unable to connect to redis db", errors.Trace())
	}
	return c, nil
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,
		// MaxActive: 12000, // max number of connections
		IdleTimeout: 120 * time.Second,

		Dial: redisConnect,

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}

			_, err := c.Do("PING")
			return err
		},
	}
}
