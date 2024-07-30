package mapsize

import (
	"math"
	_ "net/http/pprof"
	"testing"
)

func TestSize(t *testing.T) {
	m := make(map[int]struct{})
	for i := 0; i < 100; i++ {
		m[i] = struct{}{}
	}
	actual := Size(m)
	expected := int64(1416)
	if actual != expected {
		t.Fatalf("actual=%v, expected=%v", actual, expected)
	}
}

func TestPtrValueSize(t *testing.T) {
	type V struct {
		i int
	}
	m := make(map[int]*V)
	for i := 0; i < 10000; i++ {
		m[i] = &V{i: i}
	}
	actual := PtrValueSize(m)
	expected := 376664
	relError := math.Abs(1 - float64(expected)/float64(actual))
	t.Logf("relError = %f", relError)
	if relError > 0.01 {
		t.Fatalf("actual=%v, expected=%v", actual, expected)
	}
}
