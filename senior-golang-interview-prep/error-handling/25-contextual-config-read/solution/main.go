package main
import ("errors";"fmt";"io/fs";"os";"strings")
var ErrInvalidConfig=errors.New("invalid config")
func readConfig(path string)error{b,err:=os.ReadFile(path);if err!=nil{return fmt.Errorf("read config %q: %w",path,err)}
 if !strings.HasPrefix(string(b),"version="){return fmt.Errorf("read config %q: %w",path,ErrInvalidConfig)};return nil}
func main(){fmt.Println(readConfig("missing.cfg"))}
var _ error=fs.ErrNotExist

