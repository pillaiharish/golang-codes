package main
import ("errors";"fmt")
var ErrZero=errors.New("division by zero")
func Divide(a,b float64)(float64,error){if b==0{return 0,ErrZero};return a/b,nil}
func MustDivide(a,b float64)float64{v,err:=Divide(a,b);if err!=nil{panic(err)};return v}
func main(){fmt.Println(MustDivide(10,2))}

