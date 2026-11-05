package main
import ("fmt";"time")
type State int
const(Closed State=iota;Open;HalfOpen)
type Breaker struct{State State;Failures,Threshold int;OpenedAt time.Time;Cooldown time.Duration}
func(b *Breaker)Allow(now time.Time)bool{if b.State==Open&&now.Sub(b.OpenedAt)>=b.Cooldown{b.State=HalfOpen};return b.State!=Open}
func(b *Breaker)Report(now time.Time,success bool){if success{b.Failures=0;b.State=Closed;return};b.Failures++;if b.State==HalfOpen||b.Failures>=b.Threshold{b.State=Open;b.OpenedAt=now}}
func main(){b:=&Breaker{Threshold:2,Cooldown:time.Second};b.Report(time.Now(),false);fmt.Println(b.State)}

