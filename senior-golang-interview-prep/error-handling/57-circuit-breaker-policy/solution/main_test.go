package main
import ("testing";"time")
func TestBreaker(t *testing.T){now:=time.Unix(0,0);b:=&Breaker{Threshold:2,Cooldown:time.Second}
 b.Report(now,false);b.Report(now,false);if b.Allow(now)||b.State!=Open{t.Fatal("not open")}
 if !b.Allow(now.Add(time.Second))||b.State!=HalfOpen{t.Fatal("not half open")}
 b.Report(now.Add(time.Second),true);if b.State!=Closed{t.Fatal("not closed")}
}

