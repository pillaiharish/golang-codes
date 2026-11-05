package main
import ("bytes";"errors";"log/slog";"strings";"testing")
func TestStructured(t *testing.T){var b bytes.Buffer;logFailure(slog.New(slog.NewJSONHandler(&b,nil)),"u1","req1",errors.New("down"))
 for _,v:=range []string{`"user_id":"u1"`,`"request_id":"req1"`,`"error":"down"`}{if !strings.Contains(b.String(),v){t.Fatalf("missing %q in %s",v,b.String())}}
}

