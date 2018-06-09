package gache

type (
	Cache interface {
		GetValue(key string) (string, error)
		SetValue(key, value string, tags ...string) error
		Remove(key string) error
		BustTags(tags ...string) error
	}
)
