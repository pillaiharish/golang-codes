package main
import ("errors";"io/fs";"os";"path/filepath";"testing")
func TestReadConfig(t *testing.T){p:=filepath.Join(t.TempDir(),"c");if !errors.Is(readConfig(p),fs.ErrNotExist){t.Fatal("missing file not exposed")}
 if err:=os.WriteFile(p,[]byte("bad"),0600);err!=nil{t.Fatal(err)}
 if !errors.Is(readConfig(p),ErrInvalidConfig){t.Fatal("invalid config not exposed")}
 if err:=os.WriteFile(p,[]byte("version=1"),0600);err!=nil{t.Fatal(err)}
 if err:=readConfig(p);err!=nil{t.Fatal(err)}
}

