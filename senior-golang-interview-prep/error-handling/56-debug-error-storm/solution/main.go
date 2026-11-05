package main
import "fmt"
func shouldLog(n int)bool{return n==1||n%100==0}
func main(){for i:=1;i<=200;i++{if shouldLog(i){fmt.Println("sample",i)}}}

