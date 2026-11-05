package main
import ("fmt";"strconv")
func ParseAge(input string)(int,error){n,err:=strconv.Atoi(input);if err!=nil{return 0,fmt.Errorf("parse age %q: %w",input,err)};return n,nil}
func main(){fmt.Println(ParseAge("42"))}

