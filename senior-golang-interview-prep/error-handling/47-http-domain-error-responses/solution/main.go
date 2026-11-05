package main
import ("encoding/json";"errors";"net/http";"net/http/httptest";"fmt")
var ErrNotFound=errors.New("not found")
var ErrValidation=errors.New("validation")
var ErrConflict=errors.New("conflict")
var ErrUnauthorized=errors.New("unauthorized")
type apiError struct{Code string `json:"error"`;Message string `json:"message"`;RequestID string `json:"request_id"`}
func writeError(w http.ResponseWriter,err error,requestID string)error{status,code,msg:=500,"internal_error","an unexpected error occurred"
 switch{case errors.Is(err,ErrNotFound):status,code,msg=404,"not_found","user not found";case errors.Is(err,ErrValidation):status,code,msg=400,"invalid_request","invalid request";case errors.Is(err,ErrConflict):status,code,msg=409,"conflict","resource conflict";case errors.Is(err,ErrUnauthorized):status,code,msg=401,"unauthorized","authentication required"}
 w.Header().Set("Content-Type","application/json");w.WriteHeader(status);return json.NewEncoder(w).Encode(apiError{code,msg,requestID})}
func main(){w:=httptest.NewRecorder();if err:=writeError(w,fmt.Errorf("db password=secret"),"req1");err!=nil{fmt.Println(err)};fmt.Print(w.Body.String())}

