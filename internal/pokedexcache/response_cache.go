package pokedexcache

import (
	"errors"
	"sync"
	"time"
)


type CacheEntry struct {
	Val []byte
	expiredAt time.Time
}


type CacheStore struct {
	Store map[string]CacheEntry
	RwMutex *sync.RWMutex
}

func (c *CacheStore) GetCacheResponse(key string) (CacheEntry, error) {
	c.RwMutex.RLock()
	defer c.RwMutex.RUnlock()
	cacheEntry, ok := c.Store[key]
	if !ok{
		return cacheEntry,errors.New("key not found")
	}
	return cacheEntry, nil
}

func (c *CacheStore) StoreCacheEntry(key string, val []byte,expireAfter int) error {
	c.RwMutex.Lock()
	defer c.RwMutex.Unlock()
	expiredAt := time.Now().Add(time.Duration(expireAfter) * time.Second)
	cacheEntry := CacheEntry{Val:val,expiredAt: expiredAt}
	c.Store[key] = cacheEntry
	return nil
}

func (c *CacheStore) IsExpired(key string) (expired bool, err error) {
	c.RwMutex.Lock()
	defer c.RwMutex.Unlock()
	cacheEntry, ok := c.Store[key]
	if !ok{
		return false,errors.New("key not found")
	}
	currentTime := time.Now()
	if cacheEntry.expiredAt.Before(currentTime){
		return true, nil
	}
	return false, nil
}

func (c *CacheStore) RemoveCacheEntry(key string) error {
	c.RwMutex.Lock()
	defer c.RwMutex.Unlock()
	delete(c.Store,key)
	return nil
}



func CreateCacheStore() CacheStore {
	store := map[string]CacheEntry{}
	readWriteMutex := &sync.RWMutex{}

	return CacheStore{Store:store,RwMutex:readWriteMutex}
}


