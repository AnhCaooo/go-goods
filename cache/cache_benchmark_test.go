package cache

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/zap"
)

func BenchmarkCacheSet(b *testing.B) {
	logger := zap.NewNop()
	c := NewCache(logger)

	for b.Loop() {
		c.SetExpiredAfterTimePeriod("key", "value", time.Minute)
	}
}

func BenchmarkCacheGetHit(b *testing.B) {
	logger := zap.NewNop()
	c := NewCache(logger)
	c.SetExpiredAfterTimePeriod("key", "value", time.Minute)

	for b.Loop() {
		c.Get("key")
	}
}

func BenchmarkCacheGetMiss(b *testing.B) {
	logger := zap.NewNop()
	c := NewCache(logger)

	for b.Loop() {
		c.Get("unknown")
	}
}

func BenchmarkCacheGetExpired(b *testing.B) {
	logger := zap.NewNop()
	c := NewCache(logger)
	c.SetExpiredAfterTimePeriod("key", "value", -time.Second) // already expired

	for b.Loop() {
		c.Get("key")
	}
}

func BenchmarkCacheDeleteAll(b *testing.B) {
	logger := zap.NewNop()
	c := NewCache(logger)

	// Prepare 100k keys
	for i := 0; i < 100_000; i++ {
		c.SetExpiredAfterTimePeriod(fmt.Sprintf("user:%d", i), "value", time.Minute)
	}

	b.ResetTimer()

	for b.Loop() {
		c.DeleteAll("user:")
	}
}
