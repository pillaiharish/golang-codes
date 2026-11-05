package main
import ("database/sql";"errors";"testing")
func TestSQLTranslation(t *testing.T){if !errors.Is(translateGet("id",sql.ErrNoRows),ErrNotFound){t.Fatal("missing map")}
 cause:=errors.New("connection refused");if !errors.Is(translateGet("id",cause),cause){t.Fatal("lost db failure")}
 if translateGet("id",nil)!=nil{t.Fatal("nil should stay nil")}
}

