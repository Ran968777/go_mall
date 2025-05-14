package utils

import (
	"fmt"
	"sync"
	"time"
)

// Cache 缓存接口
type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, expiration time.Duration)
	Delete(key string)
}

// MemoryCache 内存缓存实现
type MemoryCache struct {
	items    map[string]cacheItem
	mu       sync.RWMutex
	stopChan chan struct{}
}

type cacheItem struct {
	value      interface{}
	expiration time.Time
}

// NewMemoryCache 创建新的内存缓存
func NewMemoryCache() *MemoryCache {
	cache := &MemoryCache{
		items:    make(map[string]cacheItem),
		stopChan: make(chan struct{}),
	}
	go cache.cleanup()
	return cache
}

// Get 获取缓存
func (c *MemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.items[key]
	if !exists {
		return nil, false
	}

	if !item.expiration.IsZero() && time.Now().After(item.expiration) {
		delete(c.items, key)
		return nil, false
	}

	return item.value, true
}

// Set 设置缓存
func (c *MemoryCache) Set(key string, value interface{}, expiration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item := cacheItem{
		value: value,
	}
	if expiration > 0 {
		item.expiration = time.Now().Add(expiration)
	}
	c.items[key] = item
}

// Delete 删除缓存
func (c *MemoryCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

// cleanup 清理过期缓存
func (c *MemoryCache) cleanup() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			now := time.Now()
			for key, item := range c.items {
				if !item.expiration.IsZero() && now.After(item.expiration) {
					delete(c.items, key)
				}
			}
			c.mu.Unlock()
		case <-c.stopChan:
			return
		}
	}
}

// Cacheable 缓存装饰器
func Cacheable(cache Cache, key string, expiration time.Duration, fn func() (interface{}, error)) (interface{}, error) {
	// 尝试从缓存获取
	if value, exists := cache.Get(key); exists {
		return value, nil
	}

	// 执行原函数
	value, err := fn()
	if err != nil {
		return nil, err
	}

	// 存入缓存
	cache.Set(key, value, expiration)
	return value, nil
}

// GenerateCacheKey 生成缓存键
func GenerateCacheKey(prefix string, params ...interface{}) string {
	key := prefix
	for _, param := range params {
		key += fmt.Sprintf(":%v", param)
	}
	return key
}
