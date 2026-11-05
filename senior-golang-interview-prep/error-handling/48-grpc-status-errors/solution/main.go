package main
import ("fmt";"google.golang.org/grpc/codes";"google.golang.org/grpc/status")
func rpcError(kind string)error{switch kind{case "missing":return status.Error(codes.NotFound,"user not found");case "invalid":return status.Error(codes.InvalidArgument,"email required");default:return status.Error(codes.Internal,"internal error")}}
func main(){fmt.Println(status.Code(rpcError("missing")))}

