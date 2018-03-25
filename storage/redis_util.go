package storage

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

const (
	rCmdExists = "EXISTS"
	rCmdSetnx  = "SETNX"
	rCmdSet    = "SET"
)

// setKV
func (r *RedisStorage) setKV(key string, v interface{}, nx bool) error {
	if nx {
		n, err := redis.Int(r.Do(rCmdExists, key))
		if err != nil {
			return err
		}
		if n == 1 {
			return fmt.Errorf("user already exists.")
		}
	}

	cmd := rCmdSet
	if nx {
		cmd = rCmdSetnx
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(v)
	if err != nil {
		return err
	}

	n, err := redis.Int(r.Do(cmd, key, buf.String()))
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("insert failed.")
	}

	return nil
}

// getKV
func (r *RedisStorage) getKV(key string, v interface{}) error {
	data, err := redis.Bytes(r.Do("GET", key))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
