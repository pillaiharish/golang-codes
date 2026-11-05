package main
import ("errors";"testing")
func TestMapping(t *testing.T){cause:=errors.New("bad");if status(&AppError{KindValidation,cause})!=400{t.Fatal("bad mapping")}
 err:=fmtError(cause);if status(err)!=503||!errors.Is(err,cause){t.Fatalf("bad transient: %v",err)}
}
func fmtError(cause error)error{return &AppError{KindTransient,cause}}

