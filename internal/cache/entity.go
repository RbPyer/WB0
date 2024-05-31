package cache

import "sync"

type Cache struct {
	sync.RWMutex
	Storage map[string]Data
}


type Data struct {
	Item interface{}
}


func NewCache() *Cache {
	return &Cache{
		Storage: make(map[string]Data),
	}
}