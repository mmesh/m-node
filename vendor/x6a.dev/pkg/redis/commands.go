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
	"github.com/gomodule/redigo/redis"
	"x6a.dev/pkg/errors"
)

func redisFlushInt(c redis.Conn) (int, error) {
	if err := c.Flush(); err != nil {
		return -1, errors.Wrapf(err, "[%v] function c.Flush()", errors.Trace())
	}

	reply, err := redis.Int(c.Receive())
	if err != nil {
		return -1, errors.Wrapf(err, "[%v] function c.Receive()", errors.Trace())
	}

	return reply, nil
}

func RedisHMSet(h string, m map[string]string) (string, error) {
	c := Pool.Get()
	defer c.Close()

	reply, err := redis.String(c.Do("HMSET", redis.Args{}.Add(h).AddFlat(m)...))
	if err != nil {
		return "", errors.Wrapf(err, "[%v] error from redis cmd HMSET hash (%v), key-values (%v)", errors.Trace(), h, m)
	}
	return reply, nil
}

func RedisHGetAll(h string) (map[string]string, error) {
	c := Pool.Get()
	defer c.Close()

	reply, err := redis.StringMap(c.Do("HGETALL", h))
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] error from redis cmd HGETALL hash (%v)", errors.Trace(), h)
	}
	return reply, nil
}

func RedisHGet(h, k string) (string, error) {
	c := Pool.Get()
	defer c.Close()

	reply, err := redis.String(c.Do("HGET", h, k))
	if err != nil {
		return "", errors.Wrapf(err, "[%v] error from redis cmd HGET hash (%v), key (%v)", errors.Trace(), h, k)
	}
	return reply, nil
}

func RedisSAdd(s, v string) (int, error) {
	c := Pool.Get()
	defer c.Close()

	reply, err := redis.Int(c.Do("SADD", s, v))
	if err != nil {
		return -1, errors.Wrapf(err, "[%v] error from redis cmd SADD set (%v), value (%v)", errors.Trace(), s, v)
	}
	return reply, nil
}

func RedisSMAdd(s string, v ...string) (int, error) {
	c := Pool.Get()
	defer c.Close()

	for _, i := range v {
		if err := c.Send("SADD", s, i); err != nil {
			return -1, errors.Wrapf(err, "[%v] error from redis cmd SREM set (%v), value (%v)", errors.Trace(), s, i)
		}
	}

	return redisFlushInt(c)
}

func RedisSRem(s, v string) (int, error) {
	c := Pool.Get()
	defer c.Close()

	reply, err := redis.Int(c.Do("SREM", s, v))
	if err != nil {
		return -1, errors.Wrapf(err, "[%v] error from redis cmd SREM set (%v), value (%v)", errors.Trace(), s, v)
	}
	return reply, nil
}

func RedisSMRem(s string, v ...string) (int, error) {
	c := Pool.Get()
	defer c.Close()

	for _, i := range v {
		if err := c.Send("SREM", s, i); err != nil {
			return -1, errors.Wrapf(err, "[%v] error from redis cmd SREM set (%v), value (%v)", errors.Trace(), s, i)
		}
	}

	return redisFlushInt(c)
}

func RedisSPop(s string) (string, error) {
	c := Pool.Get()
	defer c.Close()

	reply, err := redis.String(c.Do("SPOP", s))
	if err != nil {
		return "", errors.Wrapf(err, "[%v] error from redis cmd SPOP set (%v)", errors.Trace(), s)
	}
	if reply == "nil" {
		reply = ""
	}
	return reply, nil
}

func RedisGet(k string) (string, error) {
	c := Pool.Get()
	defer c.Close()

	reply, err := redis.String(c.Do("GET", k))
	if err != nil {
		return "", errors.Wrapf(err, "[%v] error from redis cmd GET key (%v)", errors.Trace(), k)
	}
	return reply, nil
}

func RedisSet(k, v string) (string, error) {
	c := Pool.Get()
	defer c.Close()

	reply, err := redis.String(c.Do("SET", k, v))
	if err != nil {
		return "", errors.Wrapf(err, "[%v] error from redis cmd SET key (%v), value (%v)", errors.Trace(), k, v)
	}
	return reply, nil
}

func RedisDel(k string) (int, error) {
	c := Pool.Get()
	defer c.Close()

	reply, err := redis.Int(c.Do("DEL", k))
	if err != nil {
		return -1, errors.Wrapf(err, "[%v] error from redis cmd DEL key (%v)", errors.Trace(), k)
	}
	return reply, nil
}

func RedisSMembers(s string) ([]string, error) {
	c := Pool.Get()
	defer c.Close()

	reply, err := redis.Strings(c.Do("SMEMBERS", s))
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] error from redis cmd SMEMBERS set (%v)", errors.Trace(), s)
	}
	return reply, nil
}

func RedisSISMember(s, v string) (bool, error) {
	c := Pool.Get()
	defer c.Close()

	reply, err := redis.Bool(c.Do("SISMEMBER", s, v))
	if err != nil {
		return false, errors.Wrapf(err, "[%v] error from redis cmd SISMEMBER set (%v), value (%v)", errors.Trace(), s, v)
	}
	return reply, nil
}

func RedisExists(k string) (bool, error) {
	c := Pool.Get()
	defer c.Close()

	reply, err := redis.Bool(c.Do("EXISTS", k))
	if err != nil {
		return false, errors.Wrapf(err, "[%v] error from redis cmd EXISTS key (%v)", errors.Trace(), k)
	}
	return reply, nil
}

func RedisGetSet(k, v string) (string, error) {
	c := Pool.Get()
	defer c.Close()

	reply, err := c.Do("GETSET", k, v)
	if err != nil {
		return "", errors.Wrapf(err, "[%v] error from redis cmd GETSET key (%v), value (%v)", errors.Trace(), k, v)
	}

	var r string
	switch reply.(type) {
	case []byte:
		r = string(reply.([]byte))
	default:
	}

	return r, nil
}
