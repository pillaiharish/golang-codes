package main
import ("errors";"fmt")
var ErrUnavailable=errors.New("unavailable")
type DBError struct{Code int;Err error}
func(e *DBError)Error()string{return fmt.Sprintf("db code %d: %v",e.Code,e.Err)}
func(e *DBError)Unwrap()error{return e.Err}
func query()error{return fmt.Errorf("handler: %w",fmt.Errorf("repo: %w",&DBError{503,ErrUnavailable}))}
func main(){fmt.Println(query())}

