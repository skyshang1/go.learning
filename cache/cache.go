package main

import (
	"github.com/patrickmn/go-cache"
	"time"
	"fmt"
)

func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(5 * time.Second, 10 * time.Second)

	// set delete callback
	c.OnEvicted(func(key string, value interface{}){})

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)
	c.SetDefault("baz", 43)

	// Get the string associated with the key "foo" from the cache
	for i := 0; i < 11 ; i++{
		foo, expirationTime, found := c.GetWithExpiration("foo")
		if found {
			fmt.Println(foo, expirationTime)
		} else {
			fmt.Println("can't find foo")
		}
		time.Sleep(time.Second)
	}


}
