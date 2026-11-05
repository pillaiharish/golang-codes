package main
import ("errors";"testing")
func TestValidateAll(t *testing.T){if validateAll([]string{"ok"})!=nil{t.Fatal("valid input rejected")}
 err:=validateAll([]string{"","ok",""});multi,ok:=err.(interface{Unwrap()[]error});if !ok||len(multi.Unwrap())!=2{t.Fatalf("want 2 errors: %v",err)}
 var e *ValidationError;if !errors.As(err,&e){t.Fatal("no typed validation error")}
}

