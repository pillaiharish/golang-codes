package main
import ("errors";"fmt";"sync")
var ErrA=errors.New("A")
var ErrB=errors.New("B")
func collect()error{jobs:=[]func()error{func()error{return ErrA},func()error{return nil},func()error{return ErrB}}
 ch:=make(chan error,len(jobs));var wg sync.WaitGroup
 for _,job:=range jobs{wg.Add(1);go func(f func()error){defer wg.Done();ch<-f()}(job)}
 wg.Wait();close(ch);var all []error;for err:=range ch{if err!=nil{all=append(all,err)}};return errors.Join(all...)}
func main(){fmt.Println(collect())}

