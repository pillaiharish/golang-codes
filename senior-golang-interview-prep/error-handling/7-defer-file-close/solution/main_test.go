package main
import ("errors";"os";"path/filepath";"testing")
func TestReadText(t *testing.T){p:=filepath.Join(t.TempDir(),"a");if err:=os.WriteFile(p,[]byte("hello"),0600);err!=nil{t.Fatal(err)}
 s,err:=readText(p);if err!=nil||s!="hello"{t.Fatalf("%q %v",s,err)}
 _,err=readText(p+"-missing");if !errors.Is(err,os.ErrNotExist){t.Fatalf("lost cause: %v",err)}
}

