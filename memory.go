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

func (mc *memoryCache) GetValue(key string) (string, error) {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	if elem, ok := mc.lookup[key]; ok {
		mc.items.MoveToFront(elem)
		return elem.Value.(*memoryPair).value, nil
	}

	return "", nil
}

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

func (mc *memoryCache) Remove(key string) error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	return mc.remove(key)
}

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
