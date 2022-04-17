package hw04lrucache

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
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	cacheItm := cacheItem{key: key, value: value}

	if item, ok := c.items[key]; ok {
		item.Value = cacheItm
		c.queue.MoveToFront(item)
		return true
	}

	c.queue.PushFront(cacheItm)
	c.items[key] = c.queue.Front()

	if c.queue.Len() > c.capacity {
		lastItem := c.queue.Back()
		if item, ok := lastItem.Value.(cacheItem); ok {
			delete(c.items, item.key)
		}
		c.queue.Remove(lastItem)
	}

	return false
}

func (c lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		return item.Value.(cacheItem).value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
