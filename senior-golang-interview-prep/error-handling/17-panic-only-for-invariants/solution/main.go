package main
import ("fmt";"strconv")
func ParseInt(s string)(int,error){return strconv.Atoi(s)}
func MustParseInt(s string)int{n,err:=ParseInt(s);if err!=nil{panic(fmt.Errorf("must parse integer: %w",err))};return n}
func main(){fmt.Println(MustParseInt("7"))}

