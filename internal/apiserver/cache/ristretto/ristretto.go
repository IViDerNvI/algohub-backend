package ristretto

import (
	"sync"

	"github.com/dgraph-io/ristretto/v2"
	"github.com/ividernvi/algohub/internal/apiserver/cache"
)

type datacache struct {
	cache *ristretto.Cache[string, string]
}

var (
	Once     sync.Once
	CacheIns cache.Cache
)

func GetCacheInstance() cache.Cache {
	Once.Do(func() {
		cache, err := ristretto.NewCache(&ristretto.Config[string, string]{
			NumCounters: 1e7,     // number of keys to track frequency of (10M).
			MaxCost:     1 << 30, // maximum cost of cache (1GB).
			BufferItems: 64,      // number of keys per Get buffer.
		})
		if err != nil {
			panic(err)
		}
		CacheIns = &datacache{cache: cache}
	})
	return CacheIns
}

func (d *datacache) Tokens() cache.TokenCache {
	return newTokenCache(d)
}

func (d *datacache) Close() {
	d.cache.Close()
}
