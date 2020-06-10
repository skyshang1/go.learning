package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	//c := cache.New(5*time.Second, 10*time.Second)
	//
	//// set delete callback
	//c.OnEvicted(func(key string, value interface{}) {})
	//
	//// Set the value of the key "foo" to "bar", with the default expiration time
	//c.Set("foo", "bar", cache.DefaultExpiration)
	//
	//// Set the value of the key "baz" to 42, with no expiration time
	//// (the item won't be removed until it is re-set, or removed using
	//// c.Delete("baz")
	//c.Set("baz", 42, cache.NoExpiration)
	//c.SetDefault("baz", 43)
	//
	//// Get the string associated with the key "foo" from the cache
	//for i := 0; i < 11; i++ {
	//	foo, expirationTime, found := c.GetWithExpiration("foo")
	//	if found {
	//		fmt.Println(foo, expirationTime)
	//	} else {
	//		fmt.Println("can't find foo")
	//	}
	//	time.Sleep(time.Second)
	//}

	fmt.Println("")
	c := GetInitializedCache()

	startTime := time.Now()
	// benchmark
	for index := 0; index < 10000; index++ {
		c.Get(fmt.Sprintf("index_%v", index))
	}
	fmt.Printf("Get 10000 items From Cache: %vs\n", time.Now().Sub(startTime).Seconds())

	startTime = time.Now()
	for index := 0; index < 100000; index++ {
		c.Get(fmt.Sprintf("index_%v", index))
	}
	fmt.Printf("Get 100000 items From Cache: %vs\n", time.Now().Sub(startTime).Seconds())

	startTime = time.Now()
	for index := 0; index < 1000000; index++ {
		c.Get(fmt.Sprintf("index_%v", index))
	}
	fmt.Printf("Get 1000000 items From Cache: %vs\n", time.Now().Sub(startTime).Seconds())

	startTime = time.Now()
	for index := 0; index < 10000000; index++ {
		c.Get(fmt.Sprintf("index_%v", index))
	}
	fmt.Printf("Get 10000000 items From Cache: %vs\n", time.Now().Sub(startTime).Seconds())

	startTime = time.Now()
	var mutex sync.Mutex
	for i := 0; i < 1000*1000*100; i++ {
		mutex.Lock()
		mutex.Unlock()
	}
	fmt.Printf("Execute 10000W Times Lock and Unlock Elapsed: %vs\n", time.Now().Sub(startTime).Seconds())
}

func GetInitializedCache() *cache.Cache {
	c := cache.New(5*time.Minute, 10*time.Second)

	for i := 0; i < 10000; i++ {
		c.Set(fmt.Sprintf("index_%v", i), "test", cache.DefaultExpiration)
	}

	return c
}
