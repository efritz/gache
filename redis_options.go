package gache

import "time"

// RedisConfig is a function use dot configure instances of a
// Cache which uses a Redis backend.
type RedisConfig func(*redisCache)

// WithRedisPrefix sets the prefix used in the Redis backend.
func WithRedisPrefix(prefix string) RedisConfig {
	return func(rc *redisCache) { rc.prefix = prefix }
}

// WithRedisTTL sets the TTL for keys stored in the Redis backend.
func WithRedisTTL(ttl time.Duration) RedisConfig {
	return func(rc *redisCache) { rc.ttl = int(ttl / time.Second) }
}
