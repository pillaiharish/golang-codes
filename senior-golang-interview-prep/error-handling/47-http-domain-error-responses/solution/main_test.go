package main
import ("errors";"net/http/httptest";"strings";"testing")
func TestHTTPError(t *testing.T){w:=httptest.NewRecorder();if err:=writeError(w,ErrNotFound,"abc");err!=nil{t.Fatal(err)};if w.Code!=404||!strings.Contains(w.Body.String(),`"request_id":"abc"`){t.Fatalf("%d %s",w.Code,w.Body.String())}
 w=httptest.NewRecorder();if err:=writeError(w,errors.New("sql password secret"),"id");err!=nil{t.Fatal(err)};if w.Code!=500||strings.Contains(w.Body.String(),"secret"){t.Fatalf("leak: %s",w.Body.String())}}

