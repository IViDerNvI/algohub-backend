package local

type LocalCacheOptions interface {
	NumCounters() int64
	MaxCost() int64
	BufferItems() int64
}
