package main
import ("fmt";"io";"os";"strconv")
func run(args []string,out io.Writer)error{if len(args)!=1{return fmt.Errorf("expected one integer argument")}
 n,err:=strconv.Atoi(args[0]);if err!=nil{return fmt.Errorf("parse argument: %w",err)}
 _,err=fmt.Fprintln(out,n);return err}
func exitCode(args []string,out,errs io.Writer)int{if err:=run(args,out);err!=nil{fmt.Fprintln(errs,err);return 1};return 0}
func main(){os.Exit(exitCode(os.Args[1:],os.Stdout,os.Stderr))}

