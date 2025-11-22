package cache

import (
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
)

// Cache represents a thread-safe in-memory cache with logging capabilities.
// It contains a map to store cache values, a logger for logging purposes,
// and a mutex to ensure safe concurrent access.
type Cache struct {
	Data   map[string]CacheValue
	logger *zap.Logger
	lock   sync.Mutex
}

// CacheValue represents a value stored in the cache along with its expiration time.
// Value is the actual data being cached.
// Expiration is the time at which the cached value will expire and should be considered invalid.
type CacheValue struct {
	Value      interface{}
	Expiration time.Time
}

// NewCache returns a new Cache instance
func NewCache(logger *zap.Logger) *Cache {
	return &Cache{
		Data:   make(map[string]CacheValue),
		logger: logger,
	}
}

// a method is used to add new key-value pair to the cache.
// It takes in a key, a value, and a duration representing the expiration time of the value.
// It first acquires a lock on the mutex to ensure thread safety, and then it adds the key-value pair to the map along with the expiration time.
// Finally, it releases the lock.
func (c *Cache) SetExpiredAfterTimePeriod(key string, value interface{}, duration time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()

	expirationTime := time.Now().Add(duration)
	c.logger.Debug("[go-goods] set time for cache to expired after period",
		zap.String("key", key),
		zap.Time("expired-time", expirationTime),
	)
	c.Data[key] = CacheValue{
		Value:      value,
		Expiration: expirationTime,
	}
}

// a method is used to add new key-value pair to the cache.
// It takes in a key, a value, and a time slot (by hour) representing the expiration time of the value
// It first acquires a lock on the mutex to ensure thread safety, and then it adds the key-value pair to the map along with the expiration time.
// Finally, it releases the lock.
func (c *Cache) SetExpiredAtTime(key string, value interface{}, expiredTime time.Time) {
	c.logger.Debug("[go-goods] set cache to be expired",
		zap.String("key", key),
		zap.Time("expired-time", expiredTime),
	)
	c.lock.Lock()
	defer c.lock.Unlock()

	c.Data[key] = CacheValue{
		Value:      value,
		Expiration: expiredTime,
	}
}

// a method is used to retrieve a value from the cache by using a key
// It first acquires a lock on the mutex to ensure thread safety.
// Then checks if the cache contains a value for the given key and if that value has not expired.
// If the value is still valid, it returns the value and a boolean value of `true` to indicate that a valid value was found.
// If the value is not valid (means not yet cached), it returns `nil` and a boolean value of `false`.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	value, exists := c.Data[key]
	if !exists {
		c.logger.Debug("[go-goods] cache key was not found from cache", zap.String("key", key))
		return nil, false
	}
	if time.Now().After(value.Expiration) {
		c.logger.Debug("[go-goods] cache was expired",
			zap.String("key", key),
			zap.Time("expiration-time", value.Expiration),
		)
		delete(c.Data, key)
		return nil, false
	}
	return value.Value, true
}

// Delete cache based on receiving cache key. If key is not valid, then Delete is no-op
func (c *Cache) Delete(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.Data, key)
}

// DeleteAll removes all cache entries that contain the specified key as a substring.
func (c *Cache) DeleteAll(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	for k := range c.Data {
		if strings.Contains(k, key) {
			delete(c.Data, k)
		}
	}
}
