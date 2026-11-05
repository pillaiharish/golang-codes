package main
import ("errors";"fmt";"sync")
type ValidationError struct{Index int}
func(e *ValidationError)Error()string{return fmt.Sprintf("item %d: empty",e.Index)}
func validateAll(items []string)error{type item struct{i int;s string};jobs:=make(chan item);results:=make(chan error,len(items));var wg sync.WaitGroup
 for w:=0;w<2;w++{wg.Add(1);go func(){defer wg.Done();for j:=range jobs{if j.s==""{results<-&ValidationError{j.i}}}}()}
 for i,s:=range items{jobs<-item{i,s}};close(jobs);wg.Wait();close(results);var errs []error;for e:=range results{errs=append(errs,e)};return errors.Join(errs...)}
func main(){fmt.Println(validateAll([]string{"","ok",""}))}

