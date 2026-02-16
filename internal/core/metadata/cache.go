package metadata

import (
	"sync"
	"time"

	"github.com/user/finder-clone/internal/core/fs"
)

type CacheEntry struct {
	Info       fs.FileInfo
	Expiration time.Time
}

type MetadataCache struct {
	entries map[string]CacheEntry
	mu      sync.RWMutex
	ttl     time.Duration
}

func NewMetadataCache(ttl time.Duration) *MetadataCache {
	return &MetadataCache{
		entries: make(map[string]CacheEntry),
		ttl:     ttl,
	}
}

func (c *MetadataCache) Get(path string) (fs.FileInfo, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.entries[path]
	if !ok || time.Now().After(entry.Expiration) {
		return nil, false
	}
	return entry.Info, true
}

func (c *MetadataCache) Put(path string, info fs.FileInfo) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[path] = CacheEntry{
		Info:       info,
		Expiration: time.Now().Add(c.ttl),
	}
}

func (c *MetadataCache) Invalidate(path string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.entries, path)
}
