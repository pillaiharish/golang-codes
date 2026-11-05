package main
import ("errors";"strconv";"strings";"testing")
func TestParseAge(t *testing.T){
 n,err:=ParseAge("42");if err!=nil||n!=42{t.Fatalf("%d %v",n,err)}
 _,err=ParseAge("abc");if !errors.Is(err,strconv.ErrSyntax)||!strings.Contains(err.Error(),"parse age"){t.Fatalf("missing context or cause: %v",err)}
}

