package main
import ("errors";"fmt")
type Design struct{Categories,Retries,Timeouts,Cancellation,PartialFailures,PublicMessages,Observability,Tests bool}
func Review(d Design)error{var issues []error
 for _,f:=range []struct{name string;ok bool}{{"categories",d.Categories},{"retries",d.Retries},{"timeouts",d.Timeouts},{"cancellation",d.Cancellation},{"partial failures",d.PartialFailures},{"public messages",d.PublicMessages},{"observability",d.Observability},{"tests",d.Tests}}{
  if !f.ok{issues=append(issues,fmt.Errorf("missing %s",f.name))}
 };return errors.Join(issues...)}
func main(){fmt.Println(Review(Design{}))}

