// Code generated by github.com/efritz/go-mockgen 0.1.0; DO NOT EDIT.
// This file was generated by robots at
// 2019-06-24T21:18:16-05:00
// using the command
// $ go-mockgen -f github.com/efritz/gache

package mocks

import (
	gache "github.com/efritz/gache"
	"sync"
)

// MockCache is a mock implementation of the Cache interface (from the
// package github.com/efritz/gache) used for unit testing.
type MockCache struct {
	// BustTagsFunc is an instance of a mock function object controlling the
	// behavior of the method BustTags.
	BustTagsFunc *CacheBustTagsFunc
	// GetValueFunc is an instance of a mock function object controlling the
	// behavior of the method GetValue.
	GetValueFunc *CacheGetValueFunc
	// RemoveFunc is an instance of a mock function object controlling the
	// behavior of the method Remove.
	RemoveFunc *CacheRemoveFunc
	// SetValueFunc is an instance of a mock function object controlling the
	// behavior of the method SetValue.
	SetValueFunc *CacheSetValueFunc
}

// NewMockCache creates a new mock of the Cache interface. All methods
// return zero values for all results, unless overwritten.
func NewMockCache() *MockCache {
	return &MockCache{
		BustTagsFunc: &CacheBustTagsFunc{
			defaultHook: func(...string) error {
				return nil
			},
		},
		GetValueFunc: &CacheGetValueFunc{
			defaultHook: func(string) (string, error) {
				return "", nil
			},
		},
		RemoveFunc: &CacheRemoveFunc{
			defaultHook: func(string) error {
				return nil
			},
		},
		SetValueFunc: &CacheSetValueFunc{
			defaultHook: func(string, string, ...string) error {
				return nil
			},
		},
	}
}

// NewMockCacheFrom creates a new mock of the MockCache interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockCacheFrom(i gache.Cache) *MockCache {
	return &MockCache{
		BustTagsFunc: &CacheBustTagsFunc{
			defaultHook: i.BustTags,
		},
		GetValueFunc: &CacheGetValueFunc{
			defaultHook: i.GetValue,
		},
		RemoveFunc: &CacheRemoveFunc{
			defaultHook: i.Remove,
		},
		SetValueFunc: &CacheSetValueFunc{
			defaultHook: i.SetValue,
		},
	}
}

// CacheBustTagsFunc describes the behavior when the BustTags method of the
// parent MockCache instance is invoked.
type CacheBustTagsFunc struct {
	defaultHook func(...string) error
	hooks       []func(...string) error
	history     []CacheBustTagsFuncCall
	mutex       sync.Mutex
}

// BustTags delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockCache) BustTags(v0 ...string) error {
	r0 := m.BustTagsFunc.nextHook()(v0...)
	m.BustTagsFunc.appendCall(CacheBustTagsFuncCall{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the BustTags method of
// the parent MockCache instance is invoked and the hook queue is empty.
func (f *CacheBustTagsFunc) SetDefaultHook(hook func(...string) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// BustTags method of the parent MockCache instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *CacheBustTagsFunc) PushHook(hook func(...string) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *CacheBustTagsFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(...string) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *CacheBustTagsFunc) PushReturn(r0 error) {
	f.PushHook(func(...string) error {
		return r0
	})
}

func (f *CacheBustTagsFunc) nextHook() func(...string) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *CacheBustTagsFunc) appendCall(r0 CacheBustTagsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of CacheBustTagsFuncCall objects describing
// the invocations of this function.
func (f *CacheBustTagsFunc) History() []CacheBustTagsFuncCall {
	f.mutex.Lock()
	history := make([]CacheBustTagsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// CacheBustTagsFuncCall is an object that describes an invocation of method
// BustTags on an instance of MockCache.
type CacheBustTagsFuncCall struct {
	// Arg0 is a slice containing the values of the variadic arguments
	// passed to this method invocation.
	Arg0 []string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation. The variadic slice argument is flattened in this array such
// that one positional argument and three variadic arguments would result in
// a slice of four, not two.
func (c CacheBustTagsFuncCall) Args() []interface{} {
	trailing := []interface{}{}
	for _, val := range c.Arg0 {
		trailing = append(trailing, val)
	}

	return append([]interface{}{}, trailing...)
}

// Results returns an interface slice containing the results of this
// invocation.
func (c CacheBustTagsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// CacheGetValueFunc describes the behavior when the GetValue method of the
// parent MockCache instance is invoked.
type CacheGetValueFunc struct {
	defaultHook func(string) (string, error)
	hooks       []func(string) (string, error)
	history     []CacheGetValueFuncCall
	mutex       sync.Mutex
}

// GetValue delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockCache) GetValue(v0 string) (string, error) {
	r0, r1 := m.GetValueFunc.nextHook()(v0)
	m.GetValueFunc.appendCall(CacheGetValueFuncCall{v0, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetValue method of
// the parent MockCache instance is invoked and the hook queue is empty.
func (f *CacheGetValueFunc) SetDefaultHook(hook func(string) (string, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetValue method of the parent MockCache instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *CacheGetValueFunc) PushHook(hook func(string) (string, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *CacheGetValueFunc) SetDefaultReturn(r0 string, r1 error) {
	f.SetDefaultHook(func(string) (string, error) {
		return r0, r1
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *CacheGetValueFunc) PushReturn(r0 string, r1 error) {
	f.PushHook(func(string) (string, error) {
		return r0, r1
	})
}

func (f *CacheGetValueFunc) nextHook() func(string) (string, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *CacheGetValueFunc) appendCall(r0 CacheGetValueFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of CacheGetValueFuncCall objects describing
// the invocations of this function.
func (f *CacheGetValueFunc) History() []CacheGetValueFuncCall {
	f.mutex.Lock()
	history := make([]CacheGetValueFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// CacheGetValueFuncCall is an object that describes an invocation of method
// GetValue on an instance of MockCache.
type CacheGetValueFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 string
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c CacheGetValueFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c CacheGetValueFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// CacheRemoveFunc describes the behavior when the Remove method of the
// parent MockCache instance is invoked.
type CacheRemoveFunc struct {
	defaultHook func(string) error
	hooks       []func(string) error
	history     []CacheRemoveFuncCall
	mutex       sync.Mutex
}

// Remove delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockCache) Remove(v0 string) error {
	r0 := m.RemoveFunc.nextHook()(v0)
	m.RemoveFunc.appendCall(CacheRemoveFuncCall{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Remove method of the
// parent MockCache instance is invoked and the hook queue is empty.
func (f *CacheRemoveFunc) SetDefaultHook(hook func(string) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Remove method of the parent MockCache instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *CacheRemoveFunc) PushHook(hook func(string) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *CacheRemoveFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(string) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *CacheRemoveFunc) PushReturn(r0 error) {
	f.PushHook(func(string) error {
		return r0
	})
}

func (f *CacheRemoveFunc) nextHook() func(string) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *CacheRemoveFunc) appendCall(r0 CacheRemoveFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of CacheRemoveFuncCall objects describing the
// invocations of this function.
func (f *CacheRemoveFunc) History() []CacheRemoveFuncCall {
	f.mutex.Lock()
	history := make([]CacheRemoveFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// CacheRemoveFuncCall is an object that describes an invocation of method
// Remove on an instance of MockCache.
type CacheRemoveFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c CacheRemoveFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c CacheRemoveFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// CacheSetValueFunc describes the behavior when the SetValue method of the
// parent MockCache instance is invoked.
type CacheSetValueFunc struct {
	defaultHook func(string, string, ...string) error
	hooks       []func(string, string, ...string) error
	history     []CacheSetValueFuncCall
	mutex       sync.Mutex
}

// SetValue delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockCache) SetValue(v0 string, v1 string, v2 ...string) error {
	r0 := m.SetValueFunc.nextHook()(v0, v1, v2...)
	m.SetValueFunc.appendCall(CacheSetValueFuncCall{v0, v1, v2, r0})
	return r0
}

// SetDefaultHook sets function that is called when the SetValue method of
// the parent MockCache instance is invoked and the hook queue is empty.
func (f *CacheSetValueFunc) SetDefaultHook(hook func(string, string, ...string) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// SetValue method of the parent MockCache instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *CacheSetValueFunc) PushHook(hook func(string, string, ...string) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *CacheSetValueFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(string, string, ...string) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *CacheSetValueFunc) PushReturn(r0 error) {
	f.PushHook(func(string, string, ...string) error {
		return r0
	})
}

func (f *CacheSetValueFunc) nextHook() func(string, string, ...string) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *CacheSetValueFunc) appendCall(r0 CacheSetValueFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of CacheSetValueFuncCall objects describing
// the invocations of this function.
func (f *CacheSetValueFunc) History() []CacheSetValueFuncCall {
	f.mutex.Lock()
	history := make([]CacheSetValueFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// CacheSetValueFuncCall is an object that describes an invocation of method
// SetValue on an instance of MockCache.
type CacheSetValueFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 string
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Arg2 is a slice containing the values of the variadic arguments
	// passed to this method invocation.
	Arg2 []string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation. The variadic slice argument is flattened in this array such
// that one positional argument and three variadic arguments would result in
// a slice of four, not two.
func (c CacheSetValueFuncCall) Args() []interface{} {
	trailing := []interface{}{}
	for _, val := range c.Arg2 {
		trailing = append(trailing, val)
	}

	return append([]interface{}{c.Arg0, c.Arg1}, trailing...)
}

// Results returns an interface slice containing the results of this
// invocation.
func (c CacheSetValueFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
