package main
import ("errors";"testing")
func TestFormats(t *testing.T){if errors.Is(wrapV(ErrTarget),ErrTarget)||errors.Unwrap(wrapV(ErrTarget))!=nil{t.Fatal("%v leaked identity")}
 if !errors.Is(wrapW(ErrTarget),ErrTarget){t.Fatal("%w lost identity")}
}

