package main
import ("errors";"fmt")
var ErrNotFound=errors.New("not found")
type User struct{ID int;Name string}
func FindUser(id int)(User,error){if id==1{return User{1,"Ada"},nil};return User{},ErrNotFound}
func main(){fmt.Println(FindUser(9))}

