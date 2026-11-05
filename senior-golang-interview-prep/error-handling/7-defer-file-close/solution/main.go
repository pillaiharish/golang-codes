package main
import ("fmt";"io";"os")
func readText(path string)(string,error){f,err:=os.Open(path);if err!=nil{return "",fmt.Errorf("open text %q: %w",path,err)};defer f.Close() // best-effort for read-only handle
 b,err:=io.ReadAll(f);if err!=nil{return "",fmt.Errorf("read text %q: %w",path,err)};return string(b),nil}
func main(){fmt.Println(readText("missing.txt"))}

