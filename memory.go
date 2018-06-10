package gache

import (
	"container/list"
	"sync"
)

type (
	memoryCache struct {
		items    *list.List
		lookup   map[string]*list.Element
		tags     map[string][]string
		capacity int
		mutex    sync.RWMutex
	}

	memoryPair struct {
		key   string
		value string
	}
)

// NewMemoryCache creates a Cache instance using a local memory backend.
func NewMemoryCache(configs ...MemoryConfig) Cache {
	mc := &memoryCache{
		items:    &list.List{},
		lookup:   map[string]*list.Element{},
		tags:     map[string][]string{},
		capacity: 10000,
	}

	for _, config := range configs {
		config(mc)
	}

	return mc
}

// GetValue retrieves a value associated with a key from
// memory. Accessing a key which exists will move it to the
// front of the access list.
func (mc *memoryCache) GetValue(key string) (string, error) {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	if elem, ok := mc.lookup[key]; ok {
		mc.items.MoveToFront(elem)
		return elem.Value.(*memoryPair).value, nil
	}

	return "", nil
}

// SetValue associates a key with a value in memory. If the
// insertion of an item pushes the cache over capacity, the
// key that has been accessed the least recently is removed.
func (mc *memoryCache) SetValue(key, value string, tags ...string) error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	mc.remove(key)
	mc.lookup[key] = mc.items.PushFront(&memoryPair{
		key:   key,
		value: value,
	})

	for _, tag := range tags {
		mc.tags[tag] = append(mc.tags[tag], key)
	}

	if mc.items.Len() > mc.capacity {
		mc.remove(mc.items.Back().Value.(*memoryPair).key)
	}

	return nil
}

// Remove removes a value from memory.
func (mc *memoryCache) Remove(key string) error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	return mc.remove(key)
}

// BustTags removes all keys associated with the given tags
// from memory.
func (mc *memoryCache) BustTags(tags ...string) error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	for _, tag := range tags {
		for _, key := range mc.tags[tag] {
			mc.remove(key)
		}

		delete(mc.tags, tag)
	}

	return nil
}

func (mc *memoryCache) remove(key string) error {
	if elem, ok := mc.lookup[key]; ok {
		mc.items.Remove(elem)
		delete(mc.lookup, key)
	}

	return nil
}
