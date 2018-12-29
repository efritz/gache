package gache

//go:generate go-mockgen github.com/efritz/deepjoy -i Client -d mock -f

import (
	"time"

	"github.com/aphistic/sweet"
	. "github.com/efritz/go-mockgen/matchers"
	. "github.com/onsi/gomega"

	"github.com/efritz/gache/mock"
)

type RedisSuite struct{}

func (s *RedisSuite) TestGetValue(t sweet.T) {
	client := mock.NewMockClient()
	client.ReadReplicaFunc.SetDefaultReturn(client)
	client.DoFunc.SetDefaultReturn([]uint8{'b', 'a', 'r'}, nil)

	cache := NewRedisCache(client, WithRedisPrefix("prefix"))
	Expect(cache.GetValue("foo")).To(Equal("bar"))
	Expect(client.ReadReplicaFunc).To(BeCalledOnce())
	Expect(client.DoFunc).To(BeCalledOnceWith("GET", "prefix.foo"))
}

func (s *RedisSuite) TestSetValue(t sweet.T) {
	client := mock.NewMockClient()
	cache := NewRedisCache(client, WithRedisPrefix("prefix"), WithRedisTTL(time.Second*120))
	Expect(cache.SetValue("foo", "bar", "x", "y", "z")).To(BeNil())
	Expect(client.DoFunc).To(BeCalledOnceWith(
		"EVAL",
		setScript,
		0,
		"prefix.foo",
		"bar",
		120,
		"x",
		"y",
		"z",
	))
}

func (s *RedisSuite) TestRemove(t sweet.T) {
	client := mock.NewMockClient()
	cache := NewRedisCache(client, WithRedisPrefix("prefix"), WithRedisTTL(time.Second*120))
	Expect(cache.Remove("foo")).To(BeNil())
	Expect(client.DoFunc).To(BeCalledOnceWith("DEL", "prefix.foo"))
}

func (s *RedisSuite) TestBustTags(t sweet.T) {
	client := mock.NewMockClient()
	cache := NewRedisCache(client, WithRedisPrefix("prefix"))
	Expect(cache.BustTags("x", "y", "z")).To(BeNil())
	Expect(client.DoFunc).To(BeCalledOnceWith(
		"EVAL",
		bustScript,
		0,
		"x",
		"y",
		"z",
	))
}
