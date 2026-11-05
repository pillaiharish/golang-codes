package main
import ("errors"; "testing")
func TestDivide(t *testing.T){
 for _,tc:=range []struct{a,b,want float64; fail bool}{{10,2,5,false},{-10,2,-5,false},{10,0,0,true}}{
  got,err:=Divide(tc.a,tc.b)
  if got!=tc.want || (err!=nil)!=tc.fail {t.Fatalf("%v/%v = %v, %v",tc.a,tc.b,got,err)}
  if tc.fail && !errors.Is(err,ErrDivideByZero){t.Fatalf("wrong error: %v",err)}
 }
}

