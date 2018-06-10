package gache

import (
	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

type MemorySuite struct{}

func (s *MemorySuite) TestBasicGetSet(t sweet.T) {
	c := NewMemoryCache()
	Expect(c.SetValue("k1", "v1")).To(BeNil())
	Expect(c.SetValue("k2", "v2")).To(BeNil())
	Expect(c.SetValue("k3", "v3")).To(BeNil())

	Expect(c.GetValue("k1")).To(Equal("v1"))
	Expect(c.GetValue("k2")).To(Equal("v2"))
	Expect(c.GetValue("k3")).To(Equal("v3"))
	Expect(c.GetValue("k4")).To(Equal(""))
}

func (s *MemorySuite) TestOverwriteKey(t sweet.T) {
	c := NewMemoryCache()
	Expect(c.SetValue("k1", "v1")).To(BeNil())
	Expect(c.SetValue("k1", "v2")).To(BeNil())
	Expect(c.SetValue("k1", "v3")).To(BeNil())

	v1, err := c.GetValue("k1")
	Expect(err).To(BeNil())
	Expect(v1).To(Equal("v3"))
}

func (s *MemorySuite) TestRemove(t sweet.T) {
	c := NewMemoryCache()
	Expect(c.SetValue("k1", "v1")).To(BeNil())
	Expect(c.GetValue("k1")).To(Equal("v1"))
	Expect(c.Remove("k1"))
	Expect(c.GetValue("k4")).To(Equal(""))
}

func (s *MemorySuite) TestBustTags(t sweet.T) {
	c := NewMemoryCache()
	Expect(c.SetValue("k1", "v1", "a")).To(BeNil())
	Expect(c.SetValue("k2", "v2", "b")).To(BeNil())
	Expect(c.SetValue("k3", "v3", "c")).To(BeNil())

	Expect(c.BustTags("a", "b")).To(BeNil())
	Expect(c.GetValue("k1")).To(Equal(""))
	Expect(c.GetValue("k2")).To(Equal(""))
	Expect(c.GetValue("k3")).To(Equal("v3"))
}

func (s *MemorySuite) TestBustTagsOverlappingKeys(t sweet.T) {
	c := NewMemoryCache()
	Expect(c.SetValue("k1", "v1", "a", "b")).To(BeNil())
	Expect(c.SetValue("k2", "v2", "a")).To(BeNil())
	Expect(c.SetValue("k3", "v3", "b")).To(BeNil())
	Expect(c.SetValue("k4", "v4", "c")).To(BeNil())
	Expect(c.SetValue("k5", "v5", "c")).To(BeNil())

	Expect(c.BustTags("a", "b", "c", "c")).To(BeNil())
	Expect(c.GetValue("k1")).To(Equal(""))
	Expect(c.GetValue("k2")).To(Equal(""))
	Expect(c.GetValue("k3")).To(Equal(""))
	Expect(c.GetValue("k4")).To(Equal(""))
	Expect(c.GetValue("k5")).To(Equal(""))
}

func (s *MemorySuite) TestEviction(t sweet.T) {
	c := NewMemoryCache(WithMemoryCapacity(5))
	Expect(c.SetValue("k1", "v1")).To(BeNil())
	Expect(c.SetValue("k2", "v2")).To(BeNil())
	Expect(c.SetValue("k3", "v3")).To(BeNil())
	Expect(c.SetValue("k4", "v4")).To(BeNil())
	Expect(c.SetValue("k5", "v5")).To(BeNil())

	Expect(c.GetValue("k1")).To(Equal("v1"))
	Expect(c.GetValue("k3")).To(Equal("v3"))

	Expect(c.SetValue("k6", "v6")).To(BeNil())
	Expect(c.SetValue("k7", "v7")).To(BeNil())

	Expect(c.GetValue("k1")).To(Equal("v1"))
	Expect(c.GetValue("k2")).To(Equal(""))
	Expect(c.GetValue("k3")).To(Equal("v3"))
	Expect(c.GetValue("k4")).To(Equal(""))
	Expect(c.GetValue("k5")).To(Equal("v5"))
	Expect(c.GetValue("k6")).To(Equal("v6"))
	Expect(c.GetValue("k7")).To(Equal("v7"))
}
