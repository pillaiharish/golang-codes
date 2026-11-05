package main
import ("context";"errors";"fmt";"testing")
func TestRetryClass(t *testing.T){if !IsRetryable(fmt.Errorf("op: %w",&RetryableError{errors.New("temp")})){t.Fatal("missing classification")}
 if IsRetryable(&RetryableError{context.Canceled})||IsRetryable(errors.New("bad input")){t.Fatal("misclassified")}
}

