package main
import ("bytes";"errors";"log/slog";"strings";"testing")
func TestLogOnce(t *testing.T){var b bytes.Buffer;err:=handle(slog.New(slog.NewTextHandler(&b,nil)))
 if !errors.Is(err,ErrDependency)||strings.Count(b.String(),"request failed")!=1{t.Fatalf("error=%v logs=%q",err,b.String())}
}

