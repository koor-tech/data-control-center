package cache

import "time"

const staleTime = 7 * time.Second

type CacheEntry[T any] struct {
	data      T
	UpdatedAt time.Time
	Stale     bool
}

func (c *CacheEntry[T]) Set(data T) {
	c.data = data
	c.UpdatedAt = time.Now()
	c.Stale = false
}

func (c *CacheEntry[T]) Get() (T, bool) {
	return c.data, c.IsStale()
}

func (c *CacheEntry[T]) IsStale() bool {
	return time.Now().After(c.UpdatedAt.Add(staleTime))
}
