package cache

type Cache interface {
	Tokens() TokenCache
	Close()
}

var cli *Cache

func CacheFactory() Cache {
	return *cli
}

func SetCacheFactory(cache Cache) {
	cli = &cache
}
