/**
 * @Author: kwens
 * @Date: 2022-08-29 16:27:46
 * @Description:
 */
package memcache

import (
	"sync"
	"time"
)

type MemCache struct {
	mu   sync.Mutex
	data sync.Map
}

func NewMemCache() *MemCache {
	return &MemCache{}
}

func (mem *MemCache) Set(key string, value any, opts ...MemOption) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	var defaultOpt = defaultMemOption
	for _, opt := range opts {
		opt.apply(&defaultOpt)
	}
	if defaultOpt.Timeout > 0 {
		// 设置超时时间
		go mem.setTimeout(defaultOpt.Timeout, key)
	}
	mem.data.Store(key, value)
}

func (mem *MemCache) Get(key string) any {
	mem.mu.Lock()
	defer mem.mu.Unlock()
	
	if d, ok := mem.data.Load(key); ok {
		return d
	}
	return nil
}

func (mem *MemCache) Pop(key string) any {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	if d, ok := mem.data.LoadAndDelete(key); ok {
		return d
	}
	return nil
}

func (mem *MemCache) setTimeout(d time.Duration, key string) {
	timer := time.NewTimer(d)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			mem.mu.Lock()
			mem.data.Delete(key)
			mem.mu.Unlock()
			return
		}
	}
}
