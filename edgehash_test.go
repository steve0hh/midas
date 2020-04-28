package midas

import (
	"testing"
)

func TestEdgenodeClear(t *testing.T) {
	a := NewEdgeHash(10, 4, 0)
	a.count[0][0] = 2
	a.Clear()
	if a.count[0][0] != 0 {
		t.Error("table not cleared")
	}
}
