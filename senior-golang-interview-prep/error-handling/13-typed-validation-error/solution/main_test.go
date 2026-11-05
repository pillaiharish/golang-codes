package main
import ("errors";"testing")
func TestValidation(t *testing.T){if err:=ValidateAge(1);err!=nil{t.Fatal(err)}
 var v *ValidationError;if !errors.As(ValidateAge(-1),&v)||v.Field!="age"{t.Fatalf("wrong type/field: %+v",v)}
}

