package cache

import "sync"

type Cache struct {
	sync.RWMutex
	Storage map[string]interface{}
}




func NewCache() *Cache {
	return &Cache{
		Storage: make(map[string]interface{}),
	}
}


func (c *Cache) Set(k string, v interface{}) {
	c.Lock()
	defer c.Unlock()
	c.Storage[k] = v
}

func (c *Cache) Get(k string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	record, ok := c.Storage[k]
	if !ok {
		return nil, false
	}
	return record, true
}


func (c *Cache) Del(k string) {
	c.Lock()
	defer c.Unlock()
	delete(c.Storage, k)
}
