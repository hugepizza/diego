package redis

import (
	"github.com/ckeyer/diego/storage/metadata"
	"github.com/gomodule/redigo/redis"
)

type RedisStorage struct {
	redis.Conn
}

var _ metadata.MetadataStorager = &RedisStorage{}

func NewRedisStorager(conn redis.Conn) metadata.MetadataStorager {
	return &RedisStorage{
		Conn: conn,
	}
}

const (
	keyPrefixNamespace = "ns:"
	keyPrefixUser      = "user:"
	keyPrefixOrg       = "org:"
	// keyPrefixProject   = "prj:"
)

var (
	nsKey   = redisKeyFunc(keyPrefixNamespace)
	userKey = redisKeyFunc(keyPrefixUser)
	orgKey  = redisKeyFunc(keyPrefixOrg)
)

func redisKeyFunc(prefix string) func(string) string {
	return func(name string) string {
		return prefix + name
	}
}
