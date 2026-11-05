package main
import ("database/sql";"errors";"fmt")
var ErrNotFound=errors.New("not found")
func translateGet(id string,err error)error{if err==nil{return nil};if errors.Is(err,sql.ErrNoRows){return fmt.Errorf("get user %q: %w",id,ErrNotFound)};return fmt.Errorf("get user %q: %w",id,err)}
func main(){fmt.Println(translateGet("u1",sql.ErrNoRows))}

