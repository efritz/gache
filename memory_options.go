package gache

// MemoryConfig is a function use dot configure instances of a
// Cache which uses a memory backend.
type MemoryConfig func(*memoryCache)

// WithMemoryCapacity sets the maximum number of items which are
// able to be stored in the memory backend.
func WithMemoryCapacity(capacity int) MemoryConfig {
	return func(rc *memoryCache) { rc.capacity = capacity }
}
