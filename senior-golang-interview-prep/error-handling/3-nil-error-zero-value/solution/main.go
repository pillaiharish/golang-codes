package main
import "fmt"
func successful() error { return nil }
func isNil(err error) bool { return err==nil }
func main(){var err error; fmt.Println(isNil(err),isNil(successful()))}

