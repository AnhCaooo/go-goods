package cache

import (
	"testing"
	"time"

	"go.uber.org/zap"
)

func setupTestCache(t *testing.T) *Cache {
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("failed to create logger: %v", err)
	}
	return NewCache(logger)
}

func TestNewCache(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	cache := NewCache(logger)

	if cache == nil {
		t.Error("NewCache() returned nil")
	}
	if cache.Data == nil {
		t.Error("NewCache() did not initialize Data map")
	}
	if cache.logger == nil {
		t.Error("NewCache() did not initialize logger")
	}
	if len(cache.Data) != 0 {
		t.Error("NewCache() did not initialize empty Data map")
	}
}

func TestSetExpiredAfterTimePeriod(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    interface{}
		duration time.Duration
	}{
		{
			name:     "Set with string value",
			key:      "key1",
			value:    "testValue",
			duration: 1 * time.Minute,
		},
		{
			name:     "Set with integer value",
			key:      "key2",
			value:    42,
			duration: 5 * time.Minute,
		},
		{
			name:     "Set with zero duration",
			key:      "key3",
			value:    "immediate",
			duration: 0,
		},
		{
			name:     "Set with map value",
			key:      "key4",
			value:    map[string]string{"nested": "value"},
			duration: 2 * time.Hour,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := setupTestCache(t)
			beforeSet := time.Now()
			cache.SetExpiredAfterTimePeriod(test.key, test.value, test.duration)
			afterSet := time.Now()

			if _, exists := cache.Data[test.key]; !exists {
				t.Errorf("Key %q was not added to cache", test.key)
			}

			cacheValue := cache.Data[test.key]

			// Use reflect.DeepEqual for complex types or direct comparison for primitives
			if test.value != nil {
				switch v := test.value.(type) {
				case map[string]string:
					// For maps, compare the stored map
					storedMap, ok := cacheValue.Value.(map[string]string)
					if !ok {
						t.Errorf("Expected map[string]string, got %T", cacheValue.Value)
					} else if len(storedMap) != len(v) {
						t.Errorf("Map length mismatch: expected %d, got %d", len(v), len(storedMap))
					} else {
						for key, val := range v {
							if storedMap[key] != val {
								t.Errorf("Map value mismatch for key %q: expected %v, got %v", key, val, storedMap[key])
							}
						}
					}
				default:
					// For primitive types, use direct comparison
					if cacheValue.Value != test.value {
						t.Errorf("Expected value %v, got %v", test.value, cacheValue.Value)
					}
				}
			}

			expectedMinExpiration := beforeSet.Add(test.duration)
			expectedMaxExpiration := afterSet.Add(test.duration)
			if cacheValue.Expiration.Before(expectedMinExpiration) ||
				cacheValue.Expiration.After(expectedMaxExpiration.Add(100*time.Millisecond)) {
				t.Errorf("Expiration time mismatch: got %v, expected between %v and %v",
					cacheValue.Expiration, expectedMinExpiration, expectedMaxExpiration)
			}
		})
	}
}

func TestSetExpiredAtTime(t *testing.T) {
	futureTime := time.Now().Add(1 * time.Hour)
	pastTime := time.Now().Add(-1 * time.Hour)
	currentTime := time.Now()

	tests := []struct {
		name        string
		key         string
		value       interface{}
		expiredTime time.Time
	}{
		{
			name:        "Set with future expiration",
			key:         "key1",
			value:       "testValue",
			expiredTime: futureTime,
		},
		{
			name:        "Set with past expiration",
			key:         "key2",
			value:       "expired",
			expiredTime: pastTime,
		},
		{
			name:        "Set with exact current time",
			key:         "key3",
			value:       42,
			expiredTime: currentTime,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := setupTestCache(t)
			cache.SetExpiredAtTime(test.key, test.value, test.expiredTime)

			if _, exists := cache.Data[test.key]; !exists {
				t.Errorf("Key %q was not added to cache", test.key)
			}

			cacheValue := cache.Data[test.key]
			if cacheValue.Value != test.value {
				t.Errorf("Expected value %v, got %v", test.value, cacheValue.Value)
			}

			if cacheValue.Expiration != test.expiredTime {
				t.Errorf("Expected expiration %v, got %v", test.expiredTime, cacheValue.Expiration)
			}
		})
	}
}

func TestGet(t *testing.T) {
	futureTime := time.Now().Add(1 * time.Hour)
	pastTime := time.Now().Add(-1 * time.Hour)
	nearFutureTime := time.Now().Add(30 * time.Minute)

	tests := []struct {
		name           string
		key            string
		value          interface{}
		expiration     time.Time
		expectedValue  interface{}
		expectedExists bool
	}{
		{
			name:           "Get valid non-expired value",
			key:            "key1",
			value:          "testValue",
			expiration:     futureTime,
			expectedValue:  "testValue",
			expectedExists: true,
		},
		{
			name:           "Get non-existent key",
			key:            "nonexistent",
			expectedValue:  nil,
			expectedExists: false,
		},
		{
			name:           "Get expired value",
			key:            "expiredKey",
			value:          "oldValue",
			expiration:     pastTime,
			expectedValue:  nil,
			expectedExists: false,
		},
		{
			name:           "Get with integer value",
			key:            "intKey",
			value:          123,
			expiration:     nearFutureTime,
			expectedValue:  123,
			expectedExists: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := setupTestCache(t)

			if test.key != "nonexistent" {
				cache.Data[test.key] = CacheValue{
					Value:      test.value,
					Expiration: test.expiration,
				}
			}

			value, exists := cache.Get(test.key)

			if exists != test.expectedExists {
				t.Errorf("Expected exists=%v, got %v", test.expectedExists, exists)
			}

			if exists && value != test.expectedValue {
				t.Errorf("Expected value %v, got %v", test.expectedValue, value)
			}

			// Verify expired values are deleted
			if test.expiration.Before(time.Now()) && test.key != "nonexistent" {
				if _, stillExists := cache.Data[test.key]; stillExists {
					t.Errorf("Expired key %q should have been deleted", test.key)
				}
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name        string
		keyToDelete string
		keysToAdd   []string
	}{
		{
			name:        "Delete existing key",
			keyToDelete: "key1",
			keysToAdd:   []string{"key1", "key2", "key3"},
		},
		{
			name:        "Delete non-existent key",
			keyToDelete: "nonexistent",
			keysToAdd:   []string{"key1", "key2"},
		},
		{
			name:        "Delete from single entry cache",
			keyToDelete: "only",
			keysToAdd:   []string{"only"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := setupTestCache(t)

			// Add keys to cache
			for _, key := range test.keysToAdd {
				cache.Data[key] = CacheValue{
					Value:      "value",
					Expiration: time.Now().Add(1 * time.Hour),
				}
			}

			cache.Delete(test.keyToDelete)

			// Check if key was deleted
			if _, exists := cache.Data[test.keyToDelete]; exists {
				t.Errorf("Key %q should have been deleted", test.keyToDelete)
			}

			// Verify other keys remain
			for _, key := range test.keysToAdd {
				if key != test.keyToDelete {
					if _, exists := cache.Data[key]; !exists {
						t.Errorf("Key %q should still exist", key)
					}
				}
			}
		})
	}
}

func TestDeleteAll(t *testing.T) {
	tests := []struct {
		name              string
		keysToAdd         []string
		substringToDelete string
		expectedRemaining []string
	}{
		{
			name:              "Delete all with substring match",
			keysToAdd:         []string{"user:1", "user:2", "user:3", "post:1"},
			substringToDelete: "user",
			expectedRemaining: []string{"post:1"},
		},
		{
			name:              "Delete all with no matches",
			keysToAdd:         []string{"key1", "key2", "key3"},
			substringToDelete: "nonexistent",
			expectedRemaining: []string{"key1", "key2", "key3"},
		},
		{
			name:              "Delete all entries",
			keysToAdd:         []string{"cache:1", "cache:2", "cache:3"},
			substringToDelete: "cache",
			expectedRemaining: []string{},
		},
		{
			name:              "Delete with partial substring",
			keysToAdd:         []string{"product_1", "product_2", "service_1"},
			substringToDelete: "product",
			expectedRemaining: []string{"service_1"},
		},
		{
			name:              "Delete from empty cache",
			keysToAdd:         []string{},
			substringToDelete: "any",
			expectedRemaining: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := setupTestCache(t)

			// Add keys to cache
			for _, key := range test.keysToAdd {
				cache.Data[key] = CacheValue{
					Value:      "value",
					Expiration: time.Now().Add(1 * time.Hour),
				}
			}

			cache.DeleteAll(test.substringToDelete)

			// Verify deleted keys are gone
			for _, key := range test.keysToAdd {
				if !contains(test.expectedRemaining, key) {
					if _, exists := cache.Data[key]; exists {
						t.Errorf("Key %q should have been deleted", key)
					}
				}
			}

			// Verify remaining keys still exist
			for _, key := range test.expectedRemaining {
				if _, exists := cache.Data[key]; !exists {
					t.Errorf("Key %q should still exist", key)
				}
			}

			if len(cache.Data) != len(test.expectedRemaining) {
				t.Errorf("Expected %d keys remaining, got %d", len(test.expectedRemaining), len(cache.Data))
			}
		})
	}
}

// Helper function
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
