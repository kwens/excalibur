/**
 * @Author: kwens
 * @Date: 2022-08-29 16:35:24
 * @Description:
 */
package cache

import (
	"github.com/kwens/excalibur/cache/memcache"
)

// 全局
var (
	Mem *memcache.MemCache
)

func MemInit() {
	Mem = memcache.NewMemCache()
}

func Init() {
	MemInit()
}
