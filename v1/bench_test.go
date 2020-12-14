package bench

import "testing"

func TestBench(t *testing.T) {
	if A != 1 {
		t.Fatal(A)
	}
}
