package repo

type Cache interface {
	Add(key string, value string)
	Get(key string) (string, bool)
	Len() int
}

type InMemoryCache struct {
	cache    map[string]string
	order    []string
	capacity int
}

func NewInMemoryCache(capacity int) *InMemoryCache {
	return &InMemoryCache{
		cache:    make(map[string]string),
		order:    make([]string, 0),
		capacity: capacity,
	}
}

func (c *InMemoryCache) Add(key, value string) {
	if len(c.order) >= c.capacity {
		oldest := c.order[0]
		c.order = c.order[1:]
		delete(c.cache, oldest)
	}

	c.cache[key] = value
	c.order = append(c.order, key)
}

func (c *InMemoryCache) Get(key string) (string, bool) {
	value, ok := c.cache[key]
	if !ok {
		return "", false
	}
	return value, true
}

func (c *InMemoryCache) Len() int {
	return len(c.cache)
}
