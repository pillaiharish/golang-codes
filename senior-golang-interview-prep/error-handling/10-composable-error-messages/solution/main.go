package main
import ("errors";"fmt")
var ErrConfigMissing = errors.New("config missing")
func readConfig(path string)error{return fmt.Errorf("read config %q: %w",path,ErrConfigMissing)}
func main(){fmt.Println(readConfig("app.yaml"))}

