// Package ttlcache provides a cache where items expire
package ttlcache

import (
	"context"
	"sync"
	"time"
)

type entry[V any] struct {
	expires time.Time
	value   V
}

// Cache caches values of V using keys of K. Cache operations are concurrency-safe
type Cache[K comparable, V any] struct {
	entries  map[K]entry[V]
	lock     *sync.Mutex
	lifetime time.Duration
}

// New creates a new cache that invalidates items older than lifetime, pruning
// expired items every cleanupInterval. It launches a goroutine to perform this
// cleanup which exits when ctx is cancelled.
func New[K comparable, V any](ctx context.Context, lifetime, cleanupInterval time.Duration) Cache[K, V] {
	c := Cache[K, V]{
		entries:  map[K]entry[V]{},
		lock:     &sync.Mutex{},
		lifetime: lifetime,
	}

	go func() {
		t := time.NewTicker(cleanupInterval)
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				c.Cleanup()
			}
		}
	}()
	return c
}

// Cleanup removes expired items from the cache.
func (c Cache[K, V]) Cleanup() {
	c.lock.Lock()
	defer c.lock.Unlock()
	now := time.Now()
	for k, v := range c.entries {
		if v.expires.Before(now) {
			delete(c.entries, k)
		}
	}
}

// Get a value from cache. The bool return value indicates whether or not the
// key was present.
func (c Cache[K, V]) Get(k K) (V, bool) {
	c.lock.Lock()
	v, present := c.get(k)
	c.lock.Unlock()
	return v, present
}

func (c Cache[K, V]) get(k K) (V, bool) {
	v, present := c.entries[k]
	if !present {
		return v.value, false
	}
	if v.expires.Before(time.Now()) {
		delete(c.entries, k)
		return v.value, false
	}
	return v.value, true
}

// Has checks whether a value is in the cache
func (c Cache[K, V]) Has(k K) bool {
	c.lock.Lock()
	_, present := c.get(k)
	c.lock.Unlock()
	return present
}

// Set a value in the cache
func (c Cache[K, V]) Set(k K, v V) {
	c.lock.Lock()
	c.entries[k] = entry[V]{
		expires: time.Now().Add(c.lifetime),
		value:   v,
	}
	c.lock.Unlock()
}

// Delete a value from the cache. Deleting a key that isn't present is safe
func (c Cache[K, V]) Delete(k K) {
	c.lock.Lock()
	delete(c.entries, k)
	c.lock.Unlock()
}
