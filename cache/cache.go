package cache

import (
	"github.com/metafates/geminite/config"
	"github.com/metafates/geminite/where"
	"github.com/philippgille/gokv"
	"github.com/philippgille/gokv/badgerdb"
	"github.com/philippgille/gokv/encoding"
)

type Cache[K ~string, V any] struct {
	store gokv.Store
}

func New[K ~string, V any]() (*Cache[K, V], error) {
	if !config.Config.Cache {
		return &Cache[K, V]{}, nil
	}

	store, err := badgerdb.NewStore(badgerdb.Options{
		Dir:   where.CacheDir(),
		Codec: encoding.Gob,
	})
	if err != nil {
		return nil, err
	}

	return &Cache[K, V]{
		store: store,
	}, nil
}

func (c *Cache[K, V]) Set(key K, value V) error {
	if !config.Config.Cache {
		return nil
	}

	return c.store.Set(string(key), value)
}

func (c Cache[K, V]) Get(key K) (V, bool, error) {
	var value V
	if !config.Config.Cache {
		return value, false, nil
	}

	found, err := c.store.Get(string(key), &value)

	return value, found, err
}
