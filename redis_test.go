package gache

//go:generate go-mockgen github.com/efritz/deepjoy -i Client -d mock

import (
	"time"

	"github.com/aphistic/sweet"
	"github.com/efritz/deepjoy"
	"github.com/efritz/gache/mock"
	. "github.com/onsi/gomega"
)

type RedisSuite struct{}

func (s *RedisSuite) TestGetValue(t sweet.T) {
	var (
		client = mock.NewMockClient()
		c      = NewRedisCache(
			client,
			WithRedisPrefix("prefix"),
		)
	)

	client.ReadReplicaFunc = func() deepjoy.Client {
		return client
	}

	client.DoFunc = func(command string, args ...interface{}) (interface{}, error) {
		Expect(command).To(Equal("GET"))
		Expect(args).To(HaveLen(1))
		Expect(args[0]).To(Equal("prefix.foo"))

		return []uint8{'b', 'a', 'r'}, nil
	}

	Expect(c.GetValue("foo")).To(Equal("bar"))
	Expect(client.ReadReplicaFuncCallCount()).To(Equal(1))
}

func (s *RedisSuite) TestSetValue(t sweet.T) {
	var (
		client = mock.NewMockClient()
		c      = NewRedisCache(
			client,
			WithRedisPrefix("prefix"),
			WithRedisTTL(time.Second*120),
		)
	)

	client.DoFunc = func(command string, args ...interface{}) (interface{}, error) {
		Expect(command).To(Equal("EVAL"))
		Expect(args).To(HaveLen(8))
		Expect(args[0]).To(Equal(setScript))
		Expect(args[1]).To(Equal(0))
		Expect(args[2]).To(Equal("prefix.foo"))
		Expect(args[3]).To(Equal("bar"))
		Expect(args[4]).To(Equal(120))
		Expect(args[5]).To(Equal("x"))
		Expect(args[6]).To(Equal("y"))
		Expect(args[7]).To(Equal("z"))

		return nil, nil
	}

	Expect(c.SetValue("foo", "bar", "x", "y", "z")).To(BeNil())
	Expect(client.ReadReplicaFuncCallCount()).To(Equal(0))
}

func (s *RedisSuite) TestRemove(t sweet.T) {
	var (
		client = mock.NewMockClient()
		c      = NewRedisCache(
			client,
			WithRedisPrefix("prefix"),
			WithRedisTTL(time.Second*120),
		)
	)

	client.DoFunc = func(command string, args ...interface{}) (interface{}, error) {
		Expect(command).To(Equal("DEL"))
		Expect(args).To(HaveLen(1))
		Expect(args[0]).To(Equal("prefix.foo"))

		return nil, nil
	}

	Expect(c.Remove("foo")).To(BeNil())
	Expect(client.ReadReplicaFuncCallCount()).To(Equal(0))
}

func (s *RedisSuite) TestBustTags(t sweet.T) {
	var (
		client = mock.NewMockClient()
		c      = NewRedisCache(
			client,
			WithRedisPrefix("prefix"),
		)
	)

	client.DoFunc = func(command string, args ...interface{}) (interface{}, error) {
		Expect(command).To(Equal("EVAL"))
		Expect(args).To(HaveLen(5))
		Expect(args[0]).To(Equal(bustScript))
		Expect(args[1]).To(Equal(0))
		Expect(args[2]).To(Equal("x"))
		Expect(args[3]).To(Equal("y"))
		Expect(args[4]).To(Equal("z"))

		return nil, nil
	}

	Expect(c.BustTags("x", "y", "z")).To(BeNil())
	Expect(client.ReadReplicaFuncCallCount()).To(Equal(0))
}
