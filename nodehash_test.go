package midas

import (
	"testing"
)

func TestNodehashClear(t *testing.T) {
	a := NewNodeHash(10, 2)
	a.count[0][0] = 2
	a.Clear()
	if a.count[0][0] != 0 {
		t.Error("table not cleared")
	}
}
