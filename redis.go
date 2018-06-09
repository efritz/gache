package gache

import (
	"fmt"

	"github.com/efritz/deepjoy"
	"github.com/garyburd/redigo/redis"
)

type (
	redisCache struct {
		client deepjoy.Client
		prefix string
		ttl    int
	}
)

func NewRedisCache(client deepjoy.Client, configs ...RedisConfig) Cache {
	rc := &redisCache{
		client: client,
		ttl:    3600,
	}

	for _, config := range configs {
		config(rc)
	}

	return rc
}

func (rc *redisCache) GetValue(key string) (string, error) {
	val, err := rc.client.ReadReplica().Do("GET", rc.makeKey(key))
	if err == nil && val == nil {
		return "", nil
	}

	return redis.String(val, err)
}

func (rc *redisCache) SetValue(key, value string, tags ...string) error {
	args := []interface{}{}
	args = append(args, setScript)
	args = append(args, 0)
	args = append(args, rc.makeKey(key))
	args = append(args, value)
	args = append(args, rc.ttl)
	args = append(args, tagsToArgs(tags)...)

	_, err := rc.client.Do("EVAL", args...)
	return err
}

func (rc *redisCache) Remove(key string) error {
	_, err := rc.client.Do("DEL", rc.makeKey(key))
	return err
}

func (rc *redisCache) BustTags(tags ...string) error {
	args := []interface{}{}
	args = append(args, bustScript)
	args = append(args, 0)
	args = append(args, tagsToArgs(tags)...)

	_, err := rc.client.Do("EVAL", args...)
	return err
}

func (rc *redisCache) makeKey(key string) string {
	if rc.prefix == "" {
		return key
	}

	return fmt.Sprintf("%s.%s", rc.prefix, key)
}

func tagsToArgs(tags []string) []interface{} {
	args := []interface{}{}
	for _, tag := range tags {
		args = append(args, tag)
	}

	return args
}
