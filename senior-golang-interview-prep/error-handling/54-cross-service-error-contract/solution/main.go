package main
import ("encoding/json";"fmt")
type WireError struct{Code string `json:"code"`;Message string `json:"message"`;Retryable bool `json:"retryable"`;RequestID string `json:"request_id"`}
func publicError(requestID string)([]byte,error){return json.Marshal(WireError{"USER_NOT_FOUND","user not found",false,requestID})}
func main(){b,err:=publicError("req_123");fmt.Println(string(b),err)}

