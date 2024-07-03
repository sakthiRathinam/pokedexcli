package pokedexcache

import (
	"fmt"
	"time"
)

func ReapLoop(cacheStore *CacheStore,closeChan *chan int,seconds int){
	for {
		select {
		case <- *closeChan:
			fmt.Println("closing the emitter")
			return
		default:
			time.Sleep(time.Duration(seconds) * time.Second)
			_cleanCache(cacheStore)
		}
	}
	
}

func _cleanCache(cacheStore *CacheStore){
	for key := range cacheStore.Store {
		expired, _ := cacheStore.IsExpired(key)
		if expired{
			cacheStore.RemoveCacheEntry(key)
		}
	}
}