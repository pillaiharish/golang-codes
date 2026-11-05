package main
import ("fmt";"os")
func size(path string)(int,error){data,err:=os.ReadFile(path);if err!=nil{return 0,fmt.Errorf("read input: %w",err)};return len(data),nil}
func main(){n,err:=size("missing.txt");fmt.Println(n,err)}

