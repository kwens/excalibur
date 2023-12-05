/**
 * @Author: kwens
 * @Date: 2022-11-15 16:51:47
 * @Description:
 */
package redis

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Redis interface {
	Get(key string, result interface{}) (err error)
	GetDel(key string, result interface{}) (err error)
	Set(key string, value interface{}) (err error)
	SetEx(key string, value interface{}, ex int) (err error)
	GetSet(key string, result interface{}, value interface{}) (err error)
	ExistsDel(key string) (exists bool, err error)
	Del(key string) (err error)
	Exists(key string) (exists bool, err error)
	Incr(key string) (result int, err error)
	IncrEx(key string, ex int) (result int, err error)
	TTL(key string) (result int, err error)
	LLen(key string) (result int, err error)
	RPush(key string, value interface{}) (newLen int, err error)
	RPushX(key string, value interface{}) (newLen int, err error)
	LPop(key string, result interface{}) error
	BLPop(keys []string, timeout int, result interface{}) (key string, err error)
	ScanAll(match string, count ...int) (keys []string, err error)
}

// Config 配置结构
type Config struct {
	Addr            string `json:"addr"`
	Port            int    `json:"port"`
	Password        string `json:"password"`
	DB              int    `json:"db"`
	MaxIdle         int    `json:"max_idle"`
	MaxActive       int    `json:"max_active"`
	IdleTimeout     int    `json:"idle_timeout"`
	Wait            bool   `json:"wait"`
	MaxConnLifeTime int    `json:"max_conn_life_time"`
	BaseKey         string `json:"base_key"`
}

type redisIns struct {
	config Config
	pool   *redis.Pool
}

func New(conf Config) (Redis, error) {
	_ins := &redisIns{config: conf}
	_ins.pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			addr := fmt.Sprintf("%s:%d", conf.Addr, conf.Port)
			conn, e := redis.Dial("tcp", addr, []redis.DialOption{redis.DialPassword(conf.Password),
				redis.DialDatabase(conf.DB)}...)
			if e != nil {
				return nil, e
			}
			return conn, e
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, e := c.Do("PING")
			return e
		},
		MaxIdle:         conf.MaxIdle,
		MaxActive:       conf.MaxActive,
		IdleTimeout:     time.Duration(conf.IdleTimeout),
		Wait:            conf.Wait,
		MaxConnLifetime: time.Duration(conf.MaxConnLifeTime),
	}

	return _ins, nil
}

func (r *redisIns) getKey(key string) string {
	if r.config.BaseKey != "" {
		return fmt.Sprintf("%s:%s", r.config.BaseKey, key)
	}

	return key
}

func (r *redisIns) getRawKey(key string) string {
	if r.config.BaseKey != "" {
		return strings.TrimPrefix(key, fmt.Sprintf("%s:", r.config.BaseKey))
	}

	return key
}

func (r *redisIns) getConn() redis.Conn {
	return r.pool.Get()
}

func (r *redisIns) do(fun func(c redis.Conn) (interface{}, error)) (interface{}, error) {
	conn := r.getConn()
	defer r.Close(conn)

	result, err := fun(conn)
	if err != nil {
		return nil, nil
	}
	return result, err
}
func (r *redisIns) Close(conn redis.Conn) error {
	e := conn.Close()
	if e != nil {
		log.Fatalf("conn close error :%s", e)
	}
	return e
}

func (r *redisIns) Get(key string, result interface{}) (err error) {
	reply, err := r.do(func(conn redis.Conn) (interface{}, error) {
		return conn.Do("GET", r.getKey(key))
	})
	bData, err := redis.Bytes(reply, err)
	if err != nil {
		return
	}
	err = json.Unmarshal(bData, &result)
	return err
}

func (r *redisIns) GetDel(key string, result interface{}) (err error) {
	conn := r.getConn()
	defer r.Close(conn)

	_key := r.getKey(key)
	err = conn.Send("GET", _key)
	if err != nil {
		return
	}
	err = conn.Send("DEL", _key)
	if err != nil {
		return
	}
	err = conn.Flush()
	if err != nil {
		return
	}
	bData, err := redis.Bytes(conn.Receive())
	if err != nil {
		return
	}
	_, err = conn.Receive()
	if err != nil {
		return err
	}
	err = json.Unmarshal(bData, &result)
	return err
}

func (r *redisIns) Set(key string, value interface{}) (err error) {
	data, err := json.Marshal(value)
	if err != nil {
		return
	}
	_, err = r.do(func(conn redis.Conn) (interface{}, error) {
		return conn.Do("SET", r.getKey(key), data)
	})
	return
}

func (r *redisIns) SetEx(key string, value interface{}, ex int) (err error) {
	data, err := json.Marshal(value)
	if err != nil {
		return
	}

	_, err = r.do(func(conn redis.Conn) (interface{}, error) {
		return conn.Do("SET", r.getKey(key), data, "EX", ex)
	})
	return
}

func (r *redisIns) GetSet(key string, result interface{}, value interface{}) (err error) {
	data, err := json.Marshal(value)
	if err != nil {
		return
	}

	reply, err := r.do(func(conn redis.Conn) (interface{}, error) {
		return conn.Do("GETSET", r.getKey(key), data)
	})
	bData, err := redis.Bytes(reply, err)
	if err != nil {
		return
	}
	err = json.Unmarshal(bData, &result)
	return err
}

func (r *redisIns) ExistsDel(key string) (exists bool, err error) {
	conn := r.getConn()
	defer r.Close(conn)

	_key := r.getKey(key)
	err = conn.Send("EXISTS", _key)
	if err != nil {
		return
	}
	err = conn.Send("DEL", _key)
	if err != nil {
		return
	}
	err = conn.Flush()
	if err != nil {
		return
	}
	c, err := redis.Int(conn.Receive())
	if err != nil {
		return
	}
	_, err = conn.Receive()
	if err != nil {
		return false, err
	}
	return c == 1, nil
}

func (r *redisIns) Del(key string) (err error) {
	_, err = r.do(func(conn redis.Conn) (interface{}, error) {
		return conn.Do("DEL", r.getKey(key))
	})
	return
}

func (r *redisIns) Exists(key string) (exists bool, err error) {
	reply, err := r.do(func(conn redis.Conn) (interface{}, error) {
		return conn.Do("EXISTS", r.getKey(key))
	})
	c, err := redis.Int(reply, err)
	if err != nil {
		return
	}
	return c == 1, nil
}

func (r *redisIns) Incr(key string) (result int, err error) {
	reply, err := r.do(func(conn redis.Conn) (interface{}, error) {
		return conn.Do("INCR", r.getKey(key))
	})
	res, err := redis.Int(reply, err)
	if err != nil {
		return
	}
	return res, nil
}

func (r *redisIns) IncrEx(key string, ex int) (result int, err error) {
	conn := r.getConn()
	defer r.Close(conn)

	_key := r.getKey(key)
	err = conn.Send("INCR", _key)
	if err != nil {
		return
	}
	err = conn.Send("EXPIRE", _key, ex)
	if err != nil {
		return
	}
	err = conn.Flush()
	if err != nil {
		return
	}
	res, err := redis.Int(conn.Receive())
	_, err = conn.Receive()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (r *redisIns) TTL(key string) (result int, err error) {
	reply, err := r.do(func(conn redis.Conn) (interface{}, error) {
		return conn.Do("TTL", r.getKey(key))
	})
	ttl, err := redis.Int(reply, err)
	if err != nil {
		return
	}
	return ttl, nil
}

func (r *redisIns) LLen(key string) (result int, err error) {
	reply, err := r.do(func(conn redis.Conn) (interface{}, error) {
		return conn.Do("LLEN", r.getKey(key))
	})
	ttl, err := redis.Int(reply, err)
	if err != nil {
		return
	}
	return ttl, nil
}

func (r *redisIns) RPush(key string, value interface{}) (newLen int, err error) {
	data, err := json.Marshal(value)
	if err != nil {
		return
	}
	reply, err := r.do(func(conn redis.Conn) (interface{}, error) {
		return conn.Do("RPUSH", r.getKey(key), data)
	})
	newLen, err = redis.Int(reply, err)
	return
}

func (r *redisIns) RPushX(key string, value interface{}) (newLen int, err error) {
	data, err := json.Marshal(value)
	if err != nil {
		return
	}
	reply, err := r.do(func(conn redis.Conn) (interface{}, error) {
		return conn.Do("RPUSHX", r.getKey(key), data)
	})
	newLen, err = redis.Int(reply, err)
	return
}

func (r *redisIns) LPop(key string, result interface{}) error {
	reply, err := r.do(func(conn redis.Conn) (interface{}, error) {
		return conn.Do("LPOP", r.getKey(key))
	})
	bData, err := redis.Bytes(reply, err)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bData, &result)
	if err != nil {
		return err
	}
	return nil
}

func (r *redisIns) BLPop(keys []string, timeout int, result interface{}) (key string, err error) {
	newKeys := make([]interface{}, 0)
	for _, key := range keys {
		newKeys = append(newKeys, r.getKey(key))
	}
	newKeys = append(newKeys, timeout)
	reply, err := r.do(func(conn redis.Conn) (interface{}, error) {
		return conn.Do("BLPOP", newKeys...)
	})
	strReplay, err := redis.Strings(reply, err)
	if err != nil {
		return
	}
	if len(strReplay) > 0 {
		err = json.Unmarshal([]byte(strReplay[1]), &result)
		if err != nil {
			return
		}
		return r.getRawKey(strReplay[0]), nil
	}
	return "", nil
}

func (r *redisIns) ScanAll(match string, count ...int) (keys []string, err error) {
	cursor := 0
	keys = make([]string, 0)
	for {
		pieces := make([]interface{}, 0)
		pieces = append(pieces, cursor, "MATCH", r.getKey(match))
		if len(count) > 0 {
			pieces = append(pieces, "COUNT", count[0])
		}
		reply, err := r.do(func(conn redis.Conn) (interface{}, error) {
			return conn.Do("SCAN", pieces...)
		})
		data, err := redis.Values(reply, err)
		if err != nil {
			return nil, err
		}
		cursor, err = redis.Int(data[0], nil)
		if err != nil {
			return nil, err
		}
		iterKeys, err := redis.Strings(data[1], nil)
		if err != nil {
			return nil, err
		}
		for _, k := range iterKeys {
			keys = append(keys, r.getRawKey(k))
		}
		if cursor == 0 {
			break
		}
	}
	return
}
