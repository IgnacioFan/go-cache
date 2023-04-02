package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Now")
		fk := "first key"
		sk := "second key"
		cache := NewCache()

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
	})
	http.ListenAndServe(":8080", nil)
}

type Cache struct {
	Items map[string]*Item
	mu    sync.Mutex
}

type Item struct {
	Value   string
	Expires int64
}

func NewCache() *Cache {
	cache := &Cache{Items: make(map[string]*Item)}
	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				cache.mu.Lock()
				for k, v := range cache.Items {
					if v.Expired(time.Now().UnixNano()) {
						log.Printf("%v has expires at %d", cache.Items, time.Now().UnixNano())
						delete(cache.Items, k)
					}
				}
				cache.mu.Unlock()
			}
		}
	}()
	return cache
}

func (i *Item) Expired(time int64) bool {
	if i.Expires == 0 {
		return false
	}
	return time > i.Expires
}

func (c *Cache) Get(key string) string {
	c.mu.Lock()
	var s string
	if item, ok := c.Items[key]; ok {
		s = item.Value
	}
	c.mu.Unlock()
	return s
}

func (c *Cache) Put(key, value string, expires int64) {
	c.mu.Lock()
	if _, ok := c.Items[key]; !ok {
		c.Items[key] = &Item{
			Value:   value,
			Expires: expires,
		}
	}
	c.mu.Unlock()
}
