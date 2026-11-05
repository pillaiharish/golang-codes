package main
import ("errors";"testing")
func TestJoin(t *testing.T){if ValidateUser(User{"Ada",21})!=nil{t.Fatal("valid user failed")}
 err:=ValidateUser(User{"",-1});var v *ValidationError;if !errors.As(err,&v)||v.Field!="name"{t.Fatalf("wrong first child: %+v",v)}
 multi,ok:=err.(interface{Unwrap()[]error});if !ok||len(multi.Unwrap())!=2{t.Fatalf("want two validation errors: %v",err)}
}

