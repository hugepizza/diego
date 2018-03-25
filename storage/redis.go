package storage

import (
	"time"

	"github.com/ckeyer/diego/types"
	"github.com/gomodule/redigo/redis"
)

const (
	rPreUser = "user:"
	rPreOrg  = "org:"
)

type RedisStorage struct {
	redis.Conn
}

var _ Storeger = &RedisStorage{}

func NewRedisStorager(conn redis.Conn) Storeger {
	return &RedisStorage{
		Conn: conn,
	}
}

func (r *RedisStorage) CreateUser(u *types.User) error {
	key := rPreUser + u.Name
	u.Joined = time.Now()
	return r.setKV(key, u, true)
}

// GetUser return user by name
func (r *RedisStorage) GetUser(name string) (*types.User, error) {
	key := rPreUser + name
	u := &types.User{}
	if err := r.getKV(key, u); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *RedisStorage) CreateOrg(u *types.Org) error {
	key := rPreOrg + u.Name
	return r.setKV(key, u, true)
}

// GetOrg return user by name
func (r *RedisStorage) GetOrg(name string) (*types.Org, error) {
	key := rPreOrg + name
	u := &types.Org{}
	if err := r.getKV(key, u); err != nil {
		return nil, err
	}

	return u, nil
}
