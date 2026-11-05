package main
import ("bytes";"strings";"testing")
func TestExitCode(t *testing.T){
 for _,tc:=range []struct{args []string;status int}{{[]string{"7"},0},{[]string{"bad"},1},{nil,1}}{
  var out,errs bytes.Buffer;got:=exitCode(tc.args,&out,&errs);if got!=tc.status{t.Fatalf("%v: %d",tc.args,got)}
  if got==1&&!strings.Contains(errs.String(),"argument"){t.Fatalf("no diagnostic: %q",errs.String())}
 }
}

