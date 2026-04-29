package pokecache

func (c *Cache) Get(key string) ([]byte, bool){
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.entry[key] 
	if !ok {
		return nil, false
	}
	
	return entry.val, true
}