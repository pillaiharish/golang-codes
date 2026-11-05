package main
import ("errors";"strconv";"testing")
func TestParseAge(t *testing.T){
 cases:=[]struct{name,input string;want int;wantErr bool}{{"valid","42",42,false},{"invalid","abc",0,true},{"overflow","99999999999999999999999999",0,true}}
 for _,tc:=range cases{t.Run(tc.name,func(t *testing.T){got,err:=ParseAge(tc.input)
  if got!=tc.want||(err!=nil)!=tc.wantErr{t.Fatalf("got %d,%v",got,err)}
  if tc.wantErr{var num *strconv.NumError;if !errors.As(err,&num){t.Fatalf("missing *NumError: %v",err)}}})}
}

