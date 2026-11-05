package main
import ("errors";"fmt")
type ErrorKind int
const(KindValidation ErrorKind=iota;KindNotFound;KindConflict;KindInternal;KindTransient)
type AppError struct{Kind ErrorKind;Err error}
func(e *AppError)Error()string{return fmt.Sprintf("kind %d: %v",e.Kind,e.Err)}
func(e *AppError)Unwrap()error{return e.Err}
func status(err error)int{var a *AppError;if !errors.As(err,&a){return 500};switch a.Kind{case KindValidation:return 400;case KindNotFound:return 404;case KindConflict:return 409;case KindTransient:return 503;default:return 500}}
func main(){fmt.Println(status(&AppError{KindTransient,errors.New("network")}))}

