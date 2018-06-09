package gache

type MemoryConfig func(*memoryCache)

func WithMemoryCapacity(capacity int) MemoryConfig {
	return func(rc *memoryCache) { rc.capacity = capacity }
}
