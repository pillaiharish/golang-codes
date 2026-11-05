package main
import ("bytes";"errors";"fmt";"log/slog")
var ErrDependency=errors.New("dependency down")
func service()error{return fmt.Errorf("fetch: %w",ErrDependency)}
func handle(log *slog.Logger)error{err:=service();if err!=nil{log.Error("request failed","error",err)};return err}
func main(){var b bytes.Buffer;_ = handle(slog.New(slog.NewTextHandler(&b,nil)));fmt.Print(b.String())}

