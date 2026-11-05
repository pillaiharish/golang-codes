package main
import ("errors";"fmt")
var ErrMissing=errors.New("file does not exist")
func open(path string)error{return fmt.Errorf("open %q: %w",path,ErrMissing)}
func load(path string)error{return fmt.Errorf("read config: %w",open(path))}
func main(){fmt.Println(load("app.yml"))}

