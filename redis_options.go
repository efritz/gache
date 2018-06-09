package gache

import "time"

type RedisConfig func(*redisCache)

func WithRedisPrefix(prefix string) RedisConfig {
	return func(rc *redisCache) { rc.prefix = prefix }
}

func WithRedisTTL(ttl time.Duration) RedisConfig {
	return func(rc *redisCache) { rc.ttl = int(ttl / time.Second) }
}
