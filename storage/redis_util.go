package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gomodule/redigo/redis"
)

const (
	rCmdExists = "EXISTS"
	rCmdSetnx  = "SETNX"
	rCmdSet    = "SET"
	rCmdGet    = "GET"
	rCmdMget   = "MGET"
	rCmdKeys   = "KEYS"
)

// setKV
func (r *RedisStorage) setKV(kr Keyer, nx bool) error {
	key := kr.Key()
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
	err := json.NewEncoder(buf).Encode(kr)
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

// keys
func (r *RedisStorage) getKeys(query string) ([]interface{}, error) {
	ret := []interface{}{}
	ks, err := redis.Strings(r.Do(rCmdKeys, query))
	if err != nil {
		return nil, err
	}
	for _, v := range ks {
		ret = append(ret, v)
	}
	return ret, nil
}

func (r *RedisStorage) listKVs(query string, v interface{}) error {
	ks, err := r.getKeys(query)
	if err != nil {
		return err
	}

	vstrs, err := redis.Strings(r.Do(rCmdMget, ks...))
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.WriteString("[")
	buf.WriteString(strings.Join(vstrs, ","))
	buf.WriteString("]")

	return json.NewDecoder(buf).Decode(v)
}
