package main
import "fmt"
type Kind string
const(Validation Kind="validation";NotFound Kind="not_found";Conflict Kind="conflict";RateLimit Kind="rate_limit";Dependency Kind="dependency";Internal Kind="internal")
func policy(k Kind)(status int,retryable bool){switch k{case Validation:return 400,false;case NotFound:return 404,false;case Conflict:return 409,false;case RateLimit:return 429,true;case Dependency:return 503,true;default:return 500,false}}
func main(){s,r:=policy(Dependency);fmt.Println(s,r)}

