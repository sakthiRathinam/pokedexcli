package pokedexcache

import (
	"testing"

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