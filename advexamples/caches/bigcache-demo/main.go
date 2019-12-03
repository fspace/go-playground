package main

import (
	"fmt"
	"github.com/allegro/bigcache"
	"time"
)

/**
- 从Go开源项目BigCache学习加速并发访问和避免高额的GC开销
*/
func main() {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))

	cache.Set("my-unique-key", []byte("value"))

	entry, _ := cache.Get("my-unique-key")
	fmt.Println(string(entry))
}
