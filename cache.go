package gache

type (
	// Cache is an interface to a backend key/value map with
	// additional bookkeeping for attaching tas to certain keys.
	Cache interface {
		// GetValue retrieves a value associated with a key from
		// the backend store. If the key does not exist, an empty
		// string should be returned.
		GetValue(key string) (string, error)

		// SetValue associates a key with a value in the backend
		// store. Any additional tags supplied should be attached
		// to the key. If the key already exists, the value should
		// be updated. Concrete implementations of this interface
		// are not required to clear old tags from the key (thus,
		// clients should keep tags associated with a key stable).
		SetValue(key, value string, tags ...string) error

		// Remove removes a value from the backend store.
		Remove(key string) error

		// BustTags removes all keys associated with the given tags
		// from the backend store.
		BustTags(tags ...string) error
	}
)
