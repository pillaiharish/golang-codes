package main
import ("errors";"fmt")
var ErrPermission=errors.New("permission denied")
func broken()error{return fmt.Errorf("save: %v",ErrPermission)}
func fixed()error{return fmt.Errorf("save: %w",ErrPermission)}
func main(){fmt.Println(errors.Is(broken(),ErrPermission),errors.Is(fixed(),ErrPermission))}

