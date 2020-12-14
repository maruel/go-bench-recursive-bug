package bench

import "testing"

func TestBench(t *testing.T) {
	if A != 2 {
		t.Fatal(A)
	}
}
