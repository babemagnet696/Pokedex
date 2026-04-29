package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	entry map[string]cacheEntry
	mu sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}