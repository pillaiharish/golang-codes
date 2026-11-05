package main
import ("context";"errors";"fmt")
type RetryableError struct{Err error}
func(e *RetryableError)Error()string{return "retryable: "+e.Err.Error()}
func(e *RetryableError)Unwrap()error{return e.Err}
func IsRetryable(err error)bool{if errors.Is(err,context.Canceled)||errors.Is(err,context.DeadlineExceeded){return false};var r *RetryableError;return errors.As(err,&r)}
func main(){fmt.Println(IsRetryable(&RetryableError{errors.New("temporary")}))}

