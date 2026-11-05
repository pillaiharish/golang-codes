package main
import ("errors";"testing")
func TestCompat(t *testing.T){err:=lookup();if !errors.Is(err,ErrNotFound){t.Fatal("old sentinel not recognized")};if err==ErrNotFound{t.Fatal("direct equality unexpectedly preserved")};var e *NotFoundError;if !errors.As(err,&e)||e.ID!="42"{t.Fatal("new type missing")}}

