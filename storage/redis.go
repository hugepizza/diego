package storage

import (
	"time"

	"github.com/ckeyer/diego/types"
	"github.com/ckeyer/logrus"
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
	u.Joined = time.Now()
	return r.setKV(u, true)
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

// ListUsers
func (r *RedisStorage) ListUsers() ([]*types.User, error) {
	us := []*types.User{}
	err := r.listKVs((*types.User).Prefix(nil)+"*", &us)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return us, nil
}

func (r *RedisStorage) CreateOrg(o *types.Org) error {
	o.Created = time.Now()
	return r.setKV(o, true)
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

// ListOrgs
func (r *RedisStorage) ListOrgs() ([]*types.Org, error) {
	os := []*types.Org{}
	err := r.listKVs((*types.Org).Prefix(nil)+"*", &os)
	if err != nil {
		return nil, err
	}

	return os, nil
}

// CreateProject
func (r *RedisStorage) GetProject(string) (*types.Project, error) {
	return nil, nil
}

// CreateProject
func (r *RedisStorage) CreateProject(p *types.Project) error {
	p.Created = time.Now()
	return r.setKV(p, true)
}

// CreateProject
func (r *RedisStorage) DeleteProject() {

}

// CreateFile
func (r *RedisStorage) CreateFile() {

}

// CreateFile
func (r *RedisStorage) DeleteFile() {

}
