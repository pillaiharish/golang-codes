package main
import "fmt"
func run(shouldPanic bool)(recovered any){defer func(){recovered=recover()}();if shouldPanic{panic("boom")};return nil}
func main(){fmt.Println(run(true))}

