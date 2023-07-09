package main

import "fmt"

//////////////// Interface ///////////////////////

type evictionAlgo interface {
	evict(c *cache)
}

//////////////// Strategies //////////////////////

//////////// First In First Out //////////////////
type fifo struct{}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by FIFO strategy.")
}

//////////// Least Recently Used /////////////////
type lru struct{}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by LRU strategy.")
}

//////////// Least Frequently Used ///////////////
type lfu struct{}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by LFU strategy.")
}

//////////////// Cache //////////////////////////

// cache uses a strategy
type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	fmt.Printf("Adding %s: %s\n", key, value)
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

//////////// Main ////////////////////////////////

func main() {
	lfu := &lfu{}
	cache := initCache(lfu)
	cache.add("a", "1")            // will be added to the cache
	cache.add("b", "2")            // will be added to the cache
	cache.add("c", "3")            // a value will be removed from the cache
	cache.setEvictionAlgo(&lru{})  // changing eviction algo
	cache.add("d", "4")            // a value will be removed from the cache
	cache.setEvictionAlgo(&fifo{}) // changing eviction algo
	cache.add("e", "5")            // a value will be removed from the cache
}
