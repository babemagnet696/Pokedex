package pokecache

import "time"

func NewCache(d time.Duration) *Cache {
	c := &Cache{
		entry: make(map[string]cacheEntry),
	}
	go c.reapLoop(d)
	return c
}