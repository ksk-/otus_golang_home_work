package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mtx      sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	newItem := cacheItem{key: key, value: value}

	i, ok := c.items[key]
	if ok {
		i.Value = newItem
		c.queue.MoveToFront(i)
	} else {
		if len(c.items) == c.capacity {
			removed := c.queue.Back()
			c.queue.Remove(removed)
			delete(c.items, removed.Value.(cacheItem).key)
		}

		c.items[key] = c.queue.PushFront(newItem)
	}

	return ok
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if i, ok := c.items[key]; ok {
		c.queue.MoveToFront(i)
		return i.Value.(cacheItem).value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
