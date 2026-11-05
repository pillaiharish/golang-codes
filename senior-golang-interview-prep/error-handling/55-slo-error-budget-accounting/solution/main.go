package main
import "fmt"
type SLI struct{Eligible,Success int}
func(s *SLI)Observe(status int){if status==400||status==401||status==404{return};s.Eligible++;if status<500&&status!=429{s.Success++}}
func(s SLI)Availability()float64{if s.Eligible==0{return 1};return float64(s.Success)/float64(s.Eligible)}
func main(){var s SLI;s.Observe(200);s.Observe(500);fmt.Println(s.Availability())}

