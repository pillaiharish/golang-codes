package main
import ("encoding/json";"strings";"testing")
func TestWireContract(t *testing.T){raw,err:=publicError("r1");if err!=nil{t.Fatal(err)};var got WireError
 if err:=json.Unmarshal(raw,&got);err!=nil{t.Fatal(err)};if got.Code!="USER_NOT_FOUND"||got.RequestID!="r1"||strings.Contains(string(raw),"sql"){t.Fatalf("%+v",got)}}

