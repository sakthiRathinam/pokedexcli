package pokedexcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCacheStore(t *testing.T){
	cases := []struct{
		key string
		val []byte
	}{
	{	key: "naruto",
		val: []byte("uzamaki"),
	},
	{	key: "kakashi",
		val: []byte("sensei"),
	},
	{	key: "dulquer",
		val: []byte("salmannn"),
	},

	}
	pokedexCacheStore := CreateCacheStore()
	for _, testCase := range cases {
		err := pokedexCacheStore.StoreCacheEntry(testCase.key,testCase.val,20)
		assert.Equal(t,err,nil,"Check not getting the errror")
		
		cacheVal, err := pokedexCacheStore.GetCacheResponse(testCase.key)
		assert.Equal(t,err,nil,"No error check")
		assert.Equal(t,string(cacheVal.Val),string(testCase.val),"Stored bytes check")
	}
	
}


func TestReapL(t *testing.T){
	cases := []struct {
		key string
		val []byte
		cacheExpireAfter int
	}{
	{	key: "naruto",
		val: []byte("uzamaki"),
		cacheExpireAfter: 1,
	},
	{	key: "kakashi",
		val: []byte("sensei"),
		cacheExpireAfter: 12,
	},
	{	key: "dulquer",
		val: []byte("salmannn"),
		cacheExpireAfter: 1,
	},
	}
	const reapLoopFlushInterval = 2
	pokedexCacheStore := CreateCacheStore()
	closeChan := make(chan int)
	go ReapLoop(&pokedexCacheStore,&closeChan,reapLoopFlushInterval)

	for _, testCase := range cases {
		err := pokedexCacheStore.StoreCacheEntry(testCase.key,testCase.val,testCase.cacheExpireAfter)
		assert.Equal(t,err,nil,"Check not getting the errror")
	}
	assert.Equal(t,len(pokedexCacheStore.Store),3,"checking the length before sleep")

	time.Sleep(2 * time.Second)
	assert.Equal(t,len(pokedexCacheStore.Store),1,"checking the length after sleep")
	closeChan <- 1
	defer close(closeChan)

}