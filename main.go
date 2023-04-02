package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var cache = &Cache{}

func main() {
	log.Printf("Now")

	fk := "first key"
	sk := "second key"
	cache.Put(fk, "first value", time.Now().Add(2*time.Second).UnixNano())
	s := cache.Get(fk)
	fmt.Println("Before", s)

	time.Sleep(5 * time.Second)

	s = cache.Get(fk)
	fmt.Println("After", s)

	if len(s) == 0 {
		cache.Put(sk, "second value", time.Now().Add(100*time.Second).UnixNano())
	}
	fmt.Println(cache.Get(sk))
}

type Cache struct {
	Value   sync.Map
	Expires int64
}

func (c *Cache) Expired(time int64) bool {
	if c.Expires == 0 {
		return false
	}
	return time > c.Expires
}

func (c *Cache) Get(key string) string {
	if c.Expired(time.Now().UnixNano()) {
		log.Printf("%s has expired", key)
		return ""
	}
	v, ok := c.Value.Load(key)
	var s string
	if ok {
		s, ok = v.(string)
		if !ok {
			log.Panicf("%s doesn't exist", key)
			return ""
		}
	}
	return s
}

func (c *Cache) Put(key, value string, expired int64) {
	c.Value.Store(key, value)
	c.Expires = expired
}
