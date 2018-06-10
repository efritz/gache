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

// NewRedisCache creates a Cache instance using a local memory backend.
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

// GetValue retrieves a value associated with a key from
// Redis. The read replica will be queried for this method
// only if the deepjoy Client is configured to do so.
func (rc *redisCache) GetValue(key string) (string, error) {
	val, err := rc.client.ReadReplica().Do("GET", rc.makeKey(key))
	if err == nil && val == nil {
		return "", nil
	}

	return redis.String(val, err)
}

// SetValue associates a key with a value in Redis.
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

// Remove removes a value from Redis.
func (rc *redisCache) Remove(key string) error {
	_, err := rc.client.Do("DEL", rc.makeKey(key))
	return err
}

// BustTags removes all keys associated with the given tags
// from Redis.
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
