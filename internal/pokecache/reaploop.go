package pokecache

import (
	"time"

)

func (c *Cache) reapLoop(interval time.Duration) {
	
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.entry {
			if time.Since(entry.createdAt) > interval{
				delete(c.entry, key)
			}
		}
		c.mu.Unlock()
	}
}