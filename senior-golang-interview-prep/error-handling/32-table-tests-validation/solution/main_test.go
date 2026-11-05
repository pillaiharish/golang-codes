package main
import ("errors";"strings";"testing")
func TestValidateEmail(t *testing.T){
 for _,tc:=range []struct{name,input string;bad bool}{{"valid","a@example.com",false},{"empty","",true},{"missing-at","a.example.com",true},{"long",strings.Repeat("x",255)+"@b.c",true},{"domain","a@localhost",true}}{
  t.Run(tc.name,func(t *testing.T){err:=ValidateEmail(tc.input);if (err!=nil)!=tc.bad{t.Fatalf("%v",err)}
   if tc.bad{var v *ValidationError;if !errors.Is(err,ErrInvalid)||!errors.As(err,&v)||v.Field!="email"{t.Fatalf("%v",err)}}})
 }
}

