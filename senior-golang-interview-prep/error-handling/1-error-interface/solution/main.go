package main
import "fmt"
type myError struct { msg string }
func (e *myError) Error() string { return e.msg }
func fail() error { return &myError{msg: "example failed"} }
func main() { fmt.Println(fail()) }

