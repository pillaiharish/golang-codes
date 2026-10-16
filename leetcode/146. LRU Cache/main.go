package main

import "fmt"

// Each real list node is also present in cache; head/tail are sentinels.
// The node after head is MRU; the node before tail is LRU.
type Node struct {
    key, value int
    prev, next *Node
}

type LRUCache struct {
    head, tail *Node
    cache map[int]*Node
    capacity int
}

func Constructor(capacity int) LRUCache {
    if capacity < 0 { capacity = 0 } // documented defensive extension
    cache := LRUCache{
        head: &Node{}, tail: &Node{},
        cache: make(map[int]*Node, capacity), capacity: capacity,
    }
    cache.head.next = cache.tail
    cache.tail.prev = cache.head
    return cache
}

// addToHead inserts a detached node immediately after the MRU sentinel.
func (c *LRUCache) addToHead(n *Node) {
    n.prev = c.head
    n.next = c.head.next
    c.head.next.prev = n
    c.head.next = n
}

// remove detaches a real node by rewiring both neighbors.
func (c *LRUCache) remove(n *Node) {
    n.prev.next = n.next
    n.next.prev = n.prev
    n.prev, n.next = nil, nil
}

func (c *LRUCache) Get(key int) int {
    if n, found := c.cache[key]; found {
        c.remove(n)
        c.addToHead(n)
        return n.value
    }
    return -1
}

func (c *LRUCache) Put(key, value int) {
    if n, found := c.cache[key]; found {
        n.value = value
        c.remove(n)
        c.addToHead(n)
        return
    }
    if c.capacity == 0 { return }
    n := &Node{key:key,value:value}
    c.cache[key] = n
    c.addToHead(n)
    if len(c.cache) > c.capacity {
        evicted := c.tail.prev
        c.remove(evicted)
        delete(c.cache, evicted.key)
    }
}

func main() {
    cache := Constructor(2)
    cache.Put(1, 1)
    cache.Put(2, 2)
    fmt.Println(cache.Get(1)) // 1
    cache.Put(3, 3)
    fmt.Println(cache.Get(2)) // -1
}
