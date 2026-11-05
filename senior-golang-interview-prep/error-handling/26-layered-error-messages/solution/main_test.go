package main
import ("errors";"strings";"testing")
func TestMessageLayers(t *testing.T){err:=load("app.yml");if !errors.Is(err,ErrMissing)||!strings.Contains(err.Error(),`"app.yml"`){t.Fatalf("%v",err)}}

