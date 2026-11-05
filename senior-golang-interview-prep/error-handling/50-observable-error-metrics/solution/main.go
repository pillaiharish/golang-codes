package main
import ("fmt";"sync")
type Counter struct{mu sync.Mutex;count map[string]int}
func NewCounter()*Counter{return &Counter{count:make(map[string]int)}}
func(c *Counter)Inc(kind string){c.mu.Lock();defer c.mu.Unlock();c.count[kind]++}
func(c *Counter)Get(kind string)int{c.mu.Lock();defer c.mu.Unlock();return c.count[kind]}
func main(){c:=NewCounter();c.Inc("timeout");fmt.Println(c.Get("timeout"))}

