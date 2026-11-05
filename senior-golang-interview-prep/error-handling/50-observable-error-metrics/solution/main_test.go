package main
import ("sync";"testing")
func TestCounter(t *testing.T){c:=NewCounter();var wg sync.WaitGroup;for i:=0;i<100;i++{wg.Add(1);go func(){defer wg.Done();c.Inc("timeout")}()};wg.Wait();if c.Get("timeout")!=100{t.Fatal("lost increments")}}

