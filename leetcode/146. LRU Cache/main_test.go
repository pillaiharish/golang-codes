package main

import "testing"

func TestLRUCache(t *testing.T) {
    c := Constructor(2)
    c.Put(1,1); c.Put(2,2)
    if got:=c.Get(1);got!=1 { t.Fatalf("Get(1) = %d",got) }
    c.Put(3,3) // 2 is LRU
    if got:=c.Get(2);got!=-1 { t.Fatalf("Get(2) = %d; want -1",got) }
    c.Put(4,4) // 1 is LRU
    if got:=c.Get(1);got!=-1 { t.Fatalf("Get(1) = %d; want -1",got) }
    if got:=c.Get(3);got!=3 { t.Fatalf("Get(3) = %d; want 3",got) }
    if got:=c.Get(4);got!=4 { t.Fatalf("Get(4) = %d; want 4",got) }
}

func TestLRUCacheUpdateAndCapacityZero(t *testing.T) {
    c:=Constructor(1)
    c.Put(7,10); c.Put(7,20)
    if got:=c.Get(7);got!=20 || len(c.cache)!=1 {t.Fatalf("update broke cache: value=%d, len=%d",got,len(c.cache))}
    c.Put(8,30)
    if c.Get(7)!=-1 || c.Get(8)!=30 {t.Fatal("capacity-one eviction failed")}
    zero:=Constructor(0)
    zero.Put(1,1)
    if zero.Get(1)!=-1 || len(zero.cache)!=0 {t.Fatal("zero capacity must not retain entries")}
}

func TestLRUCacheListInvariant(t *testing.T) {
    c:=Constructor(3)
    c.Put(1,10);c.Put(2,20);c.Put(3,30);c.Get(1);c.Put(4,40)
    seen:=map[*Node]bool{}
    count:=0
    for n:=c.head.next;n!=c.tail;n=n.next{
        if n==nil || n.prev==nil || n.next==nil {t.Fatal("broken list pointer")}
        if n.next.prev!=n || n.prev.next!=n {t.Fatal("broken bidirectional links")}
        if seen[n] {t.Fatal("cycle in list")}
        seen[n]=true; count++
        if c.cache[n.key]!=n {t.Fatal("list/map mismatch")}
    }
    if count!=len(c.cache)||count!=3 {t.Fatalf("list count %d; map count %d",count,len(c.cache))}
}
