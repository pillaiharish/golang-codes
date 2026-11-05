package main
import "testing"
func TestAvailability(t *testing.T){var s SLI;s.Observe(400);s.Observe(200);s.Observe(503);s.Observe(429);if s.Eligible!=3||s.Success!=1||s.Availability()!=1.0/3.0{t.Fatalf("%+v",s)}}

