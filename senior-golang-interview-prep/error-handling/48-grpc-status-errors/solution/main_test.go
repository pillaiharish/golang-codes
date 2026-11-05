package main
import ("testing";"google.golang.org/grpc/codes";"google.golang.org/grpc/status")
func TestRPCStatus(t *testing.T){for _,tc:=range []struct{kind string;want codes.Code}{{"missing",codes.NotFound},{"invalid",codes.InvalidArgument},{"other",codes.Internal}}{if got:=status.Code(rpcError(tc.kind));got!=tc.want{t.Fatalf("%q: got %v",tc.kind,got)}}}

