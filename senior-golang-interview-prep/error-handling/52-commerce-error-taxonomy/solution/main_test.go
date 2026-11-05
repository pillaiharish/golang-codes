package main
import "testing"
func TestPolicy(t *testing.T){s,r:=policy(Validation);if s!=400||r{t.Fatal("validation misclassified")};s,r=policy(Dependency);if s!=503||!r{t.Fatal("dependency misclassified")}}

