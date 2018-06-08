package gache

type (
	memoryCache struct {
		values     map[string]string
		tagMapping map[string][]string
	}
)

func NewMemoryCache() Cache {
	return &memoryCache{
		values:     map[string]string{},
		tagMapping: map[string][]string{},
	}
}

func (mc *memoryCache) GetValue(key string) (string, error) {
	if value, ok := mc.values[key]; ok {
		return value, nil
	}

	return "", nil
}

// TODO - need to make this bounded
func (mc *memoryCache) SetValue(key, value string, tags ...string) error {
	mc.values[key] = value
	for _, tag := range tags {
		mc.tagMapping[tag] = append(mc.tagMapping[tag], key)
	}

	return nil
}

func (mc *memoryCache) BustTags(tags ...string) error {
	for _, tag := range tags {
		for _, key := range mc.tagMapping[tag] {
			delete(mc.values, key)
		}

		delete(mc.tagMapping, tag)
	}

	return nil
}
