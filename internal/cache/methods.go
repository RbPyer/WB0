package cache


func (c *Cache) Set(k string, v interface{}) {
	c.Lock()
	defer c.Unlock()
	c.storage[k] = Data{Item: v}
}

func (c *Cache) Get(k string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	record, ok := c.storage[k]
	if !ok {
		return nil, false
	}
	return record.Item, true
}


func (c *Cache) Del(k string) {
	c.Lock()
	defer c.Unlock()
	delete(c.storage, k)
}