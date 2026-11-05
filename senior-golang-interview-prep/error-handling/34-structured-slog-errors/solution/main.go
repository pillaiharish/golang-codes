package main
import ("log/slog";"os")
func logFailure(logger *slog.Logger,userID,requestID string,err error){logger.Error("save user","user_id",userID,"request_id",requestID,"error",err)}
func main(){logFailure(slog.New(slog.NewJSONHandler(os.Stdout,nil)),"u1","req1",os.ErrPermission)}

