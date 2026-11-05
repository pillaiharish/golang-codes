package main
import ("errors";"testing")
func TestDBError(t *testing.T){var db *DBError;err:=query();if !errors.As(err,&db)||db.Code!=503||!errors.Is(err,ErrUnavailable){t.Fatalf("lost details: %v %+v",err,db)}}

