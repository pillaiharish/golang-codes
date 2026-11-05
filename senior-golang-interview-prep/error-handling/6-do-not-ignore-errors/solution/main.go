package main
import ("fmt";"strconv")
func unsafeParse(s string)int{n,_:=strconv.Atoi(s);return n} // lesson: do not ship
func safeParse(s string)(int,error){n,err:=strconv.Atoi(s);if err!=nil{return 0,fmt.Errorf("parse count: %w",err)};return n,nil}
func main(){fmt.Println("ignored:",unsafeParse("bad"));fmt.Println("checked:",safeParse("bad"))}

