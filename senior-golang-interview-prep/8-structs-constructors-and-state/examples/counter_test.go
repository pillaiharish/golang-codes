package main

import (
	"sync"
	"testing"
)

func TestCounterStartsEmptyAndAccumulates(t *testing.T) {
	ct := NewCounter()
	if got := ct.GetKey("missing"); got != 0 {
		t.Fatalf("missing key = %d, want 0", got)
	}
	for range 4 {
		ct.Increment("a")
	}
	ct.Increment("b")
	if got := ct.GetKey("a"); got != 4 {
		t.Fatalf("a = %d, want 4", got)
	}
	if got := ct.GetKey("b"); got != 1 {
		t.Fatalf("b = %d, want 1", got)
	}
}

func TestDifferentConstructedCountersDoNotShareState(t *testing.T) {
	first := NewCounter()
	second := NewCounter()
	first.Increment("a")
	if got := second.GetKey("a"); got != 0 {
		t.Fatalf("second counter a = %d, want 0", got)
	}
}

func TestZeroValueCounterReadButCannotWrite(t *testing.T) {
	var ct Counter // counts is nil
	if got := ct.GetKey("a"); got != 0 {
		t.Fatalf("nil-map read = %d, want 0", got)
	}
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic from assignment to a nil map")
		}
	}()
	ct.Increment("a") // Deliberate failure case, recovered by test.
}

func TestCopyOfInitializedCounterSharesMap(t *testing.T) {
	original := NewCounter()
	copyOfStruct := original // maps are descriptors; both values refer to the same map
	original.Increment("a")
	copyOfStruct.Increment("a")
	if got := original.GetKey("a"); got != 2 {
		t.Fatalf("original a = %d, want 2", got)
	}
	if got := copyOfStruct.GetKey("a"); got != 2 {
		t.Fatalf("copy a = %d, want 2", got)
	}
	// Replacing the field on one struct changes that struct's map reference.
	copyOfStruct.counts = make(map[string]int)
	copyOfStruct.Increment("a")
	if got := original.GetKey("a"); got != 2 {
		t.Fatalf("original after reassignment = %d, want 2", got)
	}
	if got := copyOfStruct.GetKey("a"); got != 1 {
		t.Fatalf("copy after reassignment = %d, want 1", got)
	}
}

func TestGaugeReceiverSemantics(t *testing.T) {
	g := Gauge{}
	g.TrySetByValue(7)
	if g.value != 0 {
		t.Fatalf("value receiver changed original to %d", g.value)
	}
	g.SetByPointer(7)
	if g.value != 7 {
		t.Fatalf("pointer receiver left original at %d", g.value)
	}
}

func TestLazyCounterZeroValueIsWritable(t *testing.T) {
	var ct LazyCounter
	ct.Increment("a")
	ct.Increment("a")
	if got := ct.GetKey("a"); got != 2 {
		t.Fatalf("lazy counter a = %d, want 2", got)
	}
}

func TestSafeCounterConcurrent(t *testing.T) {
	const workers, increments = 12, 300
	ct := NewSafeCounter()
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range increments {
				ct.Increment("shared")
			}
		}()
	}
	wg.Wait()
	if got, want := ct.GetKey("shared"), workers*increments; got != want {
		t.Fatalf("concurrent count = %d, want %d", got, want)
	}
}
