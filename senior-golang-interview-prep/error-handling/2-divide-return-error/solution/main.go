package main
import ("errors"; "fmt")
var ErrDivideByZero = errors.New("division by zero")
func Divide(a,b float64)(float64,error) { if b==0 {return 0,ErrDivideByZero}; return a/b,nil }
func main(){ n,err:=Divide(10,2); fmt.Println(n,err) }

