package main
import ("errors";"fmt")
var ErrUnavailable=errors.New("unavailable")
func oldBackend()error{return errors.New("legacy backend offline")}
func newAdapter()error{if err:=oldBackend();err!=nil{return fmt.Errorf("backend: %w",errors.Join(ErrUnavailable,err))};return nil}
func main(){fmt.Println(newAdapter())}

