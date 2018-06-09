// +build integration

package gache

import (
	"github.com/aphistic/sweet"
	"github.com/efritz/deepjoy"

	. "github.com/onsi/gomega"
)

type RedisIntegrationSuite struct{}

func RegisterIntegration(s *sweet.S) {
	// do nothing
	s.AddSuite(&RedisIntegrationSuite{})
}

func (s *RedisIntegrationSuite) SetUpSuite()    { cleanup() }
func (s *RedisIntegrationSuite) TearDownSuite() { cleanup() }

//
//

func (s *RedisIntegrationSuite) TestBasicGetSet(t sweet.T) {
	c := NewRedisCache(makeTestClient())
	Expect(c.SetValue("k1", "v1")).To(BeNil())
	Expect(c.SetValue("k2", "v2")).To(BeNil())
	Expect(c.SetValue("k3", "v3")).To(BeNil())

	Expect(c.GetValue("k1")).To(Equal("v1"))
	Expect(c.GetValue("k2")).To(Equal("v2"))
	Expect(c.GetValue("k3")).To(Equal("v3"))
	Expect(c.GetValue("k4")).To(Equal(""))
}

func (s *RedisIntegrationSuite) TestOverwriteKey(t sweet.T) {
	c := NewRedisCache(makeTestClient())
	Expect(c.SetValue("k1", "v1")).To(BeNil())
	Expect(c.SetValue("k1", "v2")).To(BeNil())
	Expect(c.SetValue("k1", "v3")).To(BeNil())

	v1, err := c.GetValue("k1")
	Expect(err).To(BeNil())
	Expect(v1).To(Equal("v3"))
}

// TODO - test remove

func (s *RedisIntegrationSuite) TestBustTags(t sweet.T) {
	c := NewRedisCache(makeTestClient())
	Expect(c.SetValue("k1", "v1", "a")).To(BeNil())
	Expect(c.SetValue("k2", "v2", "b")).To(BeNil())
	Expect(c.SetValue("k3", "v3", "c")).To(BeNil())

	Expect(c.BustTags("a", "b")).To(BeNil())
	Expect(c.GetValue("k1")).To(Equal(""))
	Expect(c.GetValue("k2")).To(Equal(""))
	Expect(c.GetValue("k3")).To(Equal("v3"))
}

func (s *RedisIntegrationSuite) TestBustTagsOverlappingKeys(t sweet.T) {
	c := NewRedisCache(makeTestClient())
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

//
//

func makeTestClient() deepjoy.Client {
	return deepjoy.NewClient("localhost:6379")
}

func cleanup() {
	makeTestClient().Do("FLUSHDB")
}
