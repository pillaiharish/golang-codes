package main
import ("errors";"strings";"testing")
func TestMessage(t *testing.T){err:=readConfig("app.yaml");if !errors.Is(err,ErrConfigMissing){t.Fatal("missing cause")}
 if strings.HasSuffix(err.Error(),".")||strings.HasPrefix(err.Error(),"Failed"){t.Fatalf("not composable: %s",err)}
}

