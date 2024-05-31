package cache

import "sync"

type Cache struct {
	sync.RWMutex
	storage map[string]Data
}


type Data struct {
	Item interface{}
}


func NewCache() *Cache {
	return &Cache{
		storage: make(map[string]Data),
	}
}