package main
import ("errors";"strings";"testing")
func TestMigration(t *testing.T){err:=newAdapter();if !errors.Is(err,ErrUnavailable)||!strings.Contains(err.Error(),"legacy backend offline"){t.Fatal(err)}}

