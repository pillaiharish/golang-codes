package main
import ("strings";"testing")
func TestReview(t *testing.T){good:=Design{true,true,true,true,true,true,true,true};if err:=Review(good);err!=nil{t.Fatal(err)}
 err:=Review(Design{Categories:true});if err==nil||!strings.Contains(err.Error(),"timeouts")||!strings.Contains(err.Error(),"tests"){t.Fatalf("%v",err)}
}

